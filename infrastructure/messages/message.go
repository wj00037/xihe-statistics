package messages

import (
	"context"
	"errors"
	"project/xihe-statistics/app"
	"project/xihe-statistics/config"
	"sync"

	"github.com/opensourceways/community-robot-lib/kafka"
	"github.com/opensourceways/community-robot-lib/mq"
	"github.com/opensourceways/community-robot-lib/utils"
	"github.com/sirupsen/logrus"
)

type message struct {
	msg  *mq.Message
	task interface{}
}

type AddRecord struct {
	hmac              string
	topic             string
	endpoint          string
	hc                utils.HttpClient
	generator         addRecordTaskGenerator
	bigModelService   app.BigModelRecordService
	repoRecordService app.RepoRecordService

	wg              sync.WaitGroup
	messageChan     chan message
	messageChanSize int
}

func NewAddRecord(cfg *config.MQConfig,
	bs app.BigModelRecordService,
	rs app.RepoRecordService,
) *AddRecord {
	size := cfg.ConcurrentSize()

	return &AddRecord{
		hmac:     cfg.AccessHmac,
		topic:    cfg.Topic,
		endpoint: cfg.AccessEndpoint,

		hc:                utils.NewHttpClient(3),
		generator:         addRecordTaskGenerator{},
		bigModelService:   bs,
		repoRecordService: rs,

		messageChan:     make(chan message, size),
		messageChanSize: size,
	}
}

func (d *AddRecord) handle(event mq.Event) error {
	msg := event.Message()

	if err := d.validateMessage(msg); err != nil {
		return err
	}

	task, ok, err := d.generator.genTask(msg.Body, msg.Header) // 处理数据，根据msg产生一个结构体数据，后面在doTask中被使用
	if err != nil || !ok {
		return err
	}

	d.messageChan <- message{
		msg:  msg,
		task: task,
	}

	return nil
}

func (d *AddRecord) validateMessage(msg *mq.Message) error {
	if msg == nil {
		return errors.New("get a nil msg from broker")
	}

	if len(msg.Header) == 0 {
		return errors.New("unexpect message: empty header")
	}

	if len(msg.Body) == 0 {
		return errors.New("unexpect message: empty payload")
	}

	return nil
}

func (d *AddRecord) doTask(log *logrus.Entry) {
	f := func(msg message) (err error) {
		task, ok := msg.task.(app.UserWithBigModelAddCmd)
		if ok {
			if err = d.bigModelService.AddUserWithBigModel(&task); err == nil { // genTask出来的数据在这里被执行
				return nil
			}
		} else {
			task, ok := msg.task.(app.RepoRecordAddCmd)
			if ok {
				if ok {
					if err = d.repoRecordService.Add(&task); err == nil { // genTask出来的数据在这里被执行
						return nil
					}
				}

				return
			}
		}
		return nil
	}

	for {
		msg, ok := <-d.messageChan
		if !ok {
			return
		}

		if err := f(msg); err != nil {
			log.Errorf("do task failed, err:%s", err.Error())
		}
	}
}

func (d *AddRecord) Run(ctx context.Context, log *logrus.Entry) error {
	s, err := kafka.Subscribe( // 订阅：创建消费组
		d.topic,
		d.handle,
		func(opt *mq.SubscribeOptions) {
			opt.Queue = "xihe-statistics" // 订阅者有同一个Queue
		},
	)
	if err != nil {
		return err
	}

	for i := 0; i < d.messageChanSize; i++ {
		d.wg.Add(1)

		go func() {
			d.doTask(log) // 这里进行了任务处理
			d.wg.Done()
		}()
	}

	<-ctx.Done()

	s.Unsubscribe()

	close(d.messageChan)

	d.wg.Wait()

	return nil
}
