package messages

import (
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Log *logrus.Entry

	MaxRetry         int
	TrainingEndpoint string

	BigModel app.BigModelRecordMessageService
	Repo     app.RepoRecordMessageService
}

func (h *Handler) AddBigModelRecord(d *domain.UserWithBigModel) error {
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
