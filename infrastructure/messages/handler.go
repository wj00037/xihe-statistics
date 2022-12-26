package messages

import (
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
)

type Handler struct {
	Log *logrus.Entry

	MaxRetry         int
	TrainingEndpoint string

	BigModel app.BigModelRecordMessageService
	Repo     app.RepoRecordMessageService
}

func (h *Handler) AddBigModelRecord(d *domain.UserWithBigModel) error { // implement domain function with app function
	cmd := app.UserWithBigModelAddCmd{
		UserWithBigModel: *d,
	}
	return h.BigModel.AddUserWithBigModel(&cmd)
}

func (h *Handler) AddRepoRecord(d *domain.UserWithRepo) error {
	cmd := app.RepoRecordAddCmd{
		UserWithRepo: *d,
	}
	return h.Repo.Add(&cmd)
}
