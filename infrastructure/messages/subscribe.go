package messages

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/opensourceways/community-robot-lib/kafka"
	"github.com/opensourceways/community-robot-lib/mq"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/config"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/message"
)

const (
	accountUnkown = "unknown"
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

	// training
	s, err := registerHandlerForTraining(handler)
	if err != nil {
		return err
	}
	if s != nil {
		subscribers[s.Topic()] = s
	}

	// register statistics
	s, err = registerHandlerForStatistics(handler)
	if err != nil {
		return err
	}
	if s != nil {
		subscribers[s.Topic()] = s
	}

	// register gitlab statistics
	s, err = registerHandlerForGitLab(handler)
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

func statisticsDo(handler interface{}, msg *mq.Message) (err error) {
	if msg == nil {
		return
	}

	body := msgStatistics{}
	if err = json.Unmarshal(msg.Body, &body); err != nil {
		return
	}

	fmt.Printf("body: %v\n", body)

	switch body.Type {
	case "bigmodel":
		h, ok := handler.(message.BigModelRecordHandler)
		if !ok {
			return
		}

		bmr := domain.UserWithBigModel{}

		if bmr.BigModel, err = domain.NewBigModel(body.Info["bigmodel"]); err != nil {
			return
		}
		if bmr.UserName, err = domain.NewAccount(body.User); err != nil {
			return
		}
		bmr.CreateAt = body.When

		return h.AddBigModelRecord(&bmr)

	case "resource":
		h, ok := handler.(message.RepoRecordHandler)
		if !ok {
			return
		}

		username, err := domain.NewAccount(body.User)
		if err != nil {
			return err
		}

		uwr := domain.UserWithRepo{
			UserName: username,
			RepoName: body.Info["name"],
			CreateAt: body.When,
		}

		return h.AddRepoRecord(&uwr)

	case "user":
		h, ok := handler.(message.RegisterRecordHandler)
		if !ok {
			return
		}

		username, err := domain.NewAccount(body.User)
		if err != nil {
			return err
		}

		d := domain.RegisterRecord{
			UserName: username,
			CreateAt: body.When,
		}

		return h.AddRegisterRecord(&d)

	case "download":
		h, ok := handler.(message.DownloadRecordHandler)
		if !ok {
			return
		}

		if body.User == "" {
			body.User = accountUnkown
		}
		username, err := domain.NewAccount(body.User)
		if err != nil {
			return err
		}

		dr := domain.DownloadRecord{
			UserName:     username,
			DownloadPath: body.Info["repo"] + "/" + body.Info["path"],
			CreateAt:     body.When,
		}

		return h.AddDownloadRecord(&dr)

	case "cloud":
		h, ok := handler.(message.CloudRecordHandler)
		if !ok {
			return
		}

		username, err := domain.NewAccount(body.User)
		if err != nil {
			return err
		}

		c := domain.Cloud{
			UserName: username,
			CloudId:  body.Info["cloud_id"],
			CreateAt: body.When,
		}

		return h.AddCloudRecord(&c)

	}

	return
}

func gitLabDo(
	handler message.FileUploadRecordHandler,
	msg *mq.Message,
) (err error) {
	if msg == nil {
		return
	}

	body := msgGitLab{}
	if err = json.Unmarshal(msg.Body, &body); err != nil {
		return
	}

	fmt.Printf("body: %v\n", body)

	if body.ObjectKind != "push" {
		return
	}

	username, err := domain.NewAccount(body.UserName)
	if err != nil {
		return
	}
	uploadPath := body.UserName + "/" + body.Project.Name
	creatAt := body.Commits[0].TimeStamp

	// tranfer time to unix time.
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	stamp, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", creatAt, local)
	if err != nil {
		return
	}
	creatAtUnix := stamp.Unix()

	fr := domain.FileUploadRecord{
		UserName:   username,
		UploadPath: uploadPath,
		CreateAt:   creatAtUnix,
	}

	return handler.AddUploadFileRecord(&fr)
}

// train
func registerHandlerForTraining(handler interface{}) (mq.Subscriber, error) {
	h, ok := handler.(message.TrainRecordHandler)
	if !ok {
		return nil, nil
	}

	return kafka.Subscribe(topics.Training, func(e mq.Event) (err error) {
		msg := e.Message()
		if msg == nil {
			return
		}

		body := MsgNormal{}
		if err = json.Unmarshal(msg.Body, &body); err != nil {
			return
		}

		if body.Details["project_id"] == "" || body.Details["training_id"] == "" {
			err = errors.New("invalid message of training")

			return
		}

		v := domain.TrainRecord{}
		if v.UserName, err = domain.NewAccount(body.User); err != nil {
			return
		}

		v.ProjectId = body.Details["project_id"]
		v.TrainId = body.Details["training_id"]
		v.CreateAt = body.CreatedAt

		return h.AddTrainRecord(&v)
	})
}

func registerHandlerForStatistics(handler interface{}) (mq.Subscriber, error) {
	return kafka.Subscribe(topics.Statistics, func(e mq.Event) (err error) {
		return statisticsDo(handler, e.Message())
	})
}

func registerHandlerForGitLab(handler interface{}) (mq.Subscriber, error) {
	h, ok := handler.(message.FileUploadRecordHandler)
	if !ok {
		return nil, nil
	}

	return kafka.Subscribe(topics.GitLab, func(e mq.Event) error {
		return gitLabDo(h, e.Message())
	})
}
