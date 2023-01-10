package messages

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

		uwr := domain.UserWithRepo{
			UserName: body.User,
			RepoName: body.Info["name"],
			CreateAt: body.When,
		}

		return h.AddRepoRecord(&uwr)

	case "user":
		h, ok := handler.(message.RegisterRecordHandler)
		if !ok {
			return
		}

		d := domain.RegisterRecord{
			UserName: body.User,
			CreateAt: body.When,
		}

		return h.AddRegisterRecord(&d)

	case "statistics-download":
		h, ok := handler.(message.DownloadRecordHandler)
		if !ok {
			return
		}

		dr := domain.DownloadRecord{
			UserName:     body.User,
			DownloadPath: body.Info["download_path"],
			CreateAt:     body.When,
		}

		return h.AddDownloadRecord(&dr)

	case "training":
		h, ok := handler.(message.TrainRecordHandler)
		if !ok {
			return
		}

		tr := domain.TrainRecord{
			UserName:  body.User,
			ProjectId: body.Info["project_id"],
			TrainId:   body.Info["id"],
			CreateAt:  body.When,
		}

		return h.AddTrainRecord(&tr)

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

	fmt.Printf("msg.Body: %v\n", msg.Body)

	body := msgGitLab{}
	if err = json.Unmarshal(msg.Body, &body); err != nil {
		return
	}

	if body.ObjectKind != "push" {
		return
	}

	fmt.Printf("body: %v\n", body)

	username := body.UserName
	uploadPath := body.UserName + "/" + body.Project.Name
	creatAt := body.Commits.TimeStamp

	// tranfer time to unix time
	local, _ := time.LoadLocation("Asia/Shanghai")
	stamp, err := time.ParseInLocation("2006-01-02T15:04:05+00:00", creatAt, local)
	if err != nil {
		return
	}
	creatAtUnix := stamp.Unix() + 8*60*60

	fr := domain.FileUploadRecord{
		UserName:   username,
		UploadPath: uploadPath,
		CreateAt:   creatAtUnix,
	}

	return handler.AddUploadFileRecord(&fr)
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
