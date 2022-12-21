package messages

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/opensourceways/community-robot-lib/kafka"
	"github.com/opensourceways/community-robot-lib/mq"
	"github.com/opensourceways/community-robot-lib/utils"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/app"
	"project/xihe-statistics/config"
)

type bigModelTask = app.UserWithBigModelAddCmd

// type repoTask = app.RepoRecordAddCmd

type msgBigModel struct {
	msg  *mq.Message
	task bigModelTask
}

type BigModel struct {
	hmac     string
	topic    string
	endpoint string
	hc       utils.HttpClient
	service  app.BigModelRecordService

	wg              sync.WaitGroup
	messageChan     chan msgBigModel
	messageChanSize int
}

func NewBigModel(cfg *config.KafKaConfig, s app.BigModelRecordService) *BigModel {
	size := cfg.ConcurrentSize()

	return &BigModel{
		hmac:     cfg.AccessHmac,
		topic:    cfg.Topic,
		endpoint: cfg.AccessEndpoint,

		hc:      utils.NewHttpClient(3),
		service: s,

		messageChan:     make(chan msgBigModel, size),
		messageChanSize: size,
	}

}

func (d *BigModel) Run(ctx context.Context, log *logrus.Entry) error {
	s, err := kafka.Subscribe(
		d.topic,
		d.handle,
		func(opt *mq.SubscribeOptions) {
			opt.Queue = "xihe-statistics"
		},
	)
	if err != nil {
		return err
	}

	for i := 0; i < d.messageChanSize; i++ {
		d.wg.Add(1)

		go func() {
			d.doTask(log)
			d.wg.Done()
		}()
	}

	<-ctx.Done()

	s.Unsubscribe()

	close(d.messageChan)

	d.wg.Wait()

	return nil
}

func (d *BigModel) handle(event mq.Event) error {
	msg := event.Message()

	if err := d.validateMessage(msg); err != nil {
		return err
	}

	cmd := new(app.UserWithBigModelAddCmd)
	err := json.Unmarshal(msg.Body, cmd)
	if err != nil {
		return err
	}

	d.messageChan <- msgBigModel{
		msg:  msg,
		task: *cmd,
	}

	return nil
}

func (d *BigModel) doTask(log *logrus.Entry) {
	f := func(msg msgBigModel) (err error) {
		task := &msg.task
		if err = d.service.AddUserWithBigModel(task); err != nil {
			return nil
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

func (d *BigModel) validateMessage(msg *mq.Message) error {
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
