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

	// register statistics
	s, err := registerHandlerForStatistics(handler)
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

func do(handler interface{}, msg *mq.Message) (err error) {
	if msg == nil {
		return
	}

	body := msgStatistics{}
	if err = json.Unmarshal(msg.Body, &body); err != nil {
		return
	}

	switch body.Type {
	case "statistics-bigmodel":
		h, ok := handler.(message.BigModelRecordHandler)
		if !ok {
			return
		}
		bmr := domain.UserWithBigModel{}
		if bmr.BigModel, err = domain.NewBigModel(body.Info); err != nil {
			return
		}

		bmr.UserName = body.UserName
		bmr.CreateAt = body.CreateAt

		return h.AddBigModelRecord(&bmr)

	case "statistics-repo":
		h, ok := handler.(message.RepoRecordHandler)
		if !ok {
			return
		}

		uwr := domain.UserWithRepo{}
		uwr.UserName = body.UserName
		uwr.RepoName = body.Info
		uwr.CreateAt = body.CreateAt

		return h.AddRepoRecord(&uwr)
	}
	return
}

func registerHandlerForStatistics(handler interface{}) (mq.Subscriber, error) {
	return kafka.Subscribe(topics.Statistics, func(e mq.Event) (err error) {
		return do(handler, e.Message())
	})
}
