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

	BigModel   app.BigModelRecordService
	Repo       app.RepoRecordService
	Register   app.RegisterRecordService
	FileUpload app.FileUploadRecordService
	Download   app.DownloadRecordService
	Train      app.TrainRecordService
	Cloud      app.CloudRecordService
}

func (h *Handler) AddBigModelRecord(d *domain.UserWithBigModel) error { // implement domain function with app function
	cmd := app.UserWithBigModelAddCmd{
		UserName: d.UserName,
		BigModel: d.BigModel,
		CreatAt:  d.CreateAt,
	}

	return h.BigModel.AddUserWithBigModel(&cmd)
}

func (h *Handler) AddRepoRecord(d *domain.UserWithRepo) error {
	cmd := app.RepoRecordAddCmd{
		UserWithRepo: *d,
	}

	if err := cmd.Validate(); err != nil {
		return err
	}

	return h.Repo.Add(&cmd)
}

func (h *Handler) AddRegisterRecord(d *domain.RegisterRecord) error {
	cmd := app.RegisterRecordAddCmd{
		RegisterRecord: *d,
	}

	if err := cmd.Validate(); err != nil {
		return err
	}

	return h.Register.Add(&cmd)
}

func (h *Handler) AddUploadFileRecord(d *domain.FileUploadRecord) error {
	cmd := app.FileUploadRecordAddCmd{
		FileUploadRecord: *d,
	}

	if err := cmd.Validate(); err != nil {
		return err
	}

	return h.FileUpload.AddRecord(cmd)
}

func (h *Handler) AddDownloadRecord(d *domain.DownloadRecord) error {
	cmd := app.DownloadRecordAddCmd{
		DownloadRecord: *d,
	}

	if err := cmd.Validate(); err != nil {
		return err
	}

	return h.Download.Add(&cmd)
}

// training
func (h *Handler) AddTrainRecord(d *domain.TrainRecord) error {
	cmd := app.TrainRecordAddCmd{
		TrainRecord: *d,
	}

	if err := cmd.Validate(); err != nil {
		return err
	}

	return h.Train.Add(&cmd)
}

func (h *Handler) AddCloudRecord(d *domain.Cloud) error {
	cmd := app.CloudRecordCmd{
		User:     d.UserName,
		CloudId:  d.CloudId,
		CreateAt: d.CreateAt,
	}

	if err := cmd.Validate(); err != nil {
		return err
	}

	return h.Cloud.Add(&cmd)
}
