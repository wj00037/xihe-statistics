package messages

import (
	"context"
	"encoding/json"
	"errors"

	kafka "github.com/opensourceways/kafka-lib/agent"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/config"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/message"
	"project/xihe-statistics/utils"
)

const (
	accountUnkown = "unknown"
	group         = "xihe-statistics"
	retry         = 3

	bigModelType = "bigmodel_type"
)

var topics config.Topics

func Subscribe(ctx context.Context, handler interface{}, log *logrus.Entry) error {
	// training
	err := registerHandlerForTraining(handler)
	if err != nil {
		return err
	}

	// register statistics
	err = registerHandlerForStatistics(handler)
	if err != nil {
		return err
	}

	// register gitlab statistics
	err = registerHandlerForGitLab(handler)
	if err != nil {
		return err
	}

	// register bigmdoel access
	if err = registerHandlerForBigModelStarted(handler); err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}

func statisticsDo(handler interface{}, b []byte) (err error) {
	if len(b) == 0 {
		return
	}

	body := msgStatistics{}
	if err = json.Unmarshal(b, &body); err != nil {
		return
	}

	switch body.Type {
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
	b []byte,
) (err error) {
	if len(b) == 0 {
		return
	}

	body := msgGitLab{}
	if err = json.Unmarshal(b, &body); err != nil {
		return
	}

	if body.ObjectKind != "push" {
		return
	}

	username, err := domain.NewAccount(body.UserName)
	if err != nil {
		return
	}
	uploadPath := body.UserName + "/" + body.Project.Name
	if len(body.Commits) < 1 {
		return
	}

	createdAtUnix, err := utils.TimeStampToUnixTime(body.Commits[0].TimeStamp)
	if err != nil {
		return
	}

	fr := domain.FileUploadRecord{
		UserName:   username,
		UploadPath: uploadPath,
		CreateAt:   createdAtUnix,
	}

	return handler.AddUploadFileRecord(&fr)
}

func subscribe(topic string, handler kafka.Handler) error {
	return kafka.SubscribeWithStrategyOfRetry(group, handler, []string{topic}, retry)
}

// train
func registerHandlerForTraining(handler interface{}) error {
	h, ok := handler.(message.TrainRecordHandler)
	if !ok {
		return nil
	}

	return subscribe(topics.Training, func(b []byte, m map[string]string) (err error) {
		if len(b) == 0 {
			return
		}

		body := MsgNormal{}
		if err = json.Unmarshal(b, &body); err != nil {
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

func registerHandlerForStatistics(handler interface{}) error {
	err := subscribe(topics.Statistics, func(b []byte, m map[string]string) (err error) {
		return statisticsDo(handler, b)
	})
	if err != nil {
		return err
	}

	return subscribe(topics.Cloud, func(b []byte, m map[string]string) (err error) {
		return statisticsDo(handler, b)
	})
}

func registerHandlerForGitLab(handler interface{}) error {
	h, ok := handler.(message.FileUploadRecordHandler)
	if !ok {
		return nil
	}

	return subscribe(topics.GitLab, func(b []byte, m map[string]string) error {
		return gitLabDo(h, b)
	})
}

func registerHandlerForBigModelStarted(handler interface{}) error {
	h, ok := handler.(message.BigModelRecordHandler)
	if !ok {
		return errors.New("handler assert error")
	}

	return subscribe(topics.BigModelStarted, func(b []byte, m map[string]string) (err error) {
		body := MsgNormal{}
		if err = json.Unmarshal(b, &body); err != nil {
			return
		}

		bmr := domain.UserWithBigModel{}

		if bmr.BigModel, err = domain.NewBigModel(body.Details[bigModelType]); err != nil {
			return
		}
		if bmr.UserName, err = domain.NewAccount(body.User); err != nil {
			return
		}
		bmr.CreateAt = body.CreatedAt

		return h.AddBigModelRecord(&bmr)
	})
}
