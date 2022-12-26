package messages

import (
	"context"
	"encoding/json"

	"github.com/opensourceways/community-robot-lib/kafka"
	"github.com/opensourceways/community-robot-lib/mq"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/config"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/message"
)

var topics config.Topics

func Subscribe(ctx context.Context, handler interface{}, log *logrus.Entry) error {
	subscribers := make(map[string]mq.Subscriber)

	defer func() {
		for k, s := range subscribers {
			if err := s.Unsubscribe(); err != nil {
				log.Errorf("failed to unsubscribe for topic:%s, err:%v", k, err)
			}
		}
	}()

	// register bigmodel
	s, err := registerHandlerForBigModelRecord(handler)
	if err != nil {
		return err
	}
	if s != nil {
		subscribers[s.Topic()] = s
	}

	// register repo
	s, err = registerHandlerForRepoRecord(handler)
	if err != nil {
		return err
	}
	if s != nil {
		subscribers[s.Topic()] = s
	}

	// register end
	if len(subscribers) == 0 {
		return nil
	}

	<-ctx.Done()

	return nil
}

func registerHandlerForBigModelRecord(handler interface{}) (mq.Subscriber, error) {
	h, ok := handler.(message.BigModelRecordHandler)
	if !ok {
		return nil, nil
	}

	return kafka.Subscribe(topics.StatisticsBigModel, func(e mq.Event) (err error) {
		msg := e.Message()
		if msg == nil {
			return
		}

		body := msgBigModelRecord{}
		if err = json.Unmarshal(msg.Body, &body); err != nil {
			return
		}
		bmr := domain.UserWithBigModel{}
		if bmr.BigModel, err = domain.NewBigModel(body.BigModel); err != nil {
			return
		}

		bmr.UserName = body.UserName
		bmr.CreateAt = body.CreateAt

		return h.AddBigModelRecord(&bmr)
	})
}

func registerHandlerForRepoRecord(handler interface{}) (mq.Subscriber, error) {
	h, ok := handler.(message.RepoRecordHandler)
	if !ok {
		return nil, nil
	}

	return kafka.Subscribe(topics.StatisticsRepo, func(e mq.Event) (err error) {
		msg := e.Message()
		if msg == nil {
			return
		}

		body := msgRepoRecord{}
		if err = json.Unmarshal(msg.Body, &body); err != nil {
			return
		}

		uwr := domain.UserWithRepo{}
		uwr.UserName = body.UserName
		uwr.RepoName = body.RepoName
		uwr.CreateAt = body.CreateAt

		return h.AddRepoRecord(&uwr)
	})
}
