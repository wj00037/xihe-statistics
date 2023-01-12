package app

import (
	"errors"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type FileUploadRecordService interface {
	GetUsersCounts() (FileUploadRecordDTO, error)
	AddRecord(FileUploadRecordAddCmd) error
}

func NewFileUploadRecordService(
	fr repository.FileUploadRecord,
) FileUploadRecordService {
	return fileUploadRecordService{
		fr: fr,
	}
}

type fileUploadRecordService struct {
	fr repository.FileUploadRecord
}

func (s fileUploadRecordService) GetUsersCounts() (
	dto FileUploadRecordDTO,
	err error,
) {
	fu, err := s.fr.Get()
	users := fu.Users
	dto = FileUploadRecordDTO{
		Users:    users,
		Counts:   int64(len(users)),
		UpdateAt: getLocalTime(),
	}

	return
}

func (s fileUploadRecordService) AddRecord(
	cmd FileUploadRecordAddCmd,
) (err error) {
	d := new(domain.FileUploadRecord)

	cmd.toFileUploadRecord(d)
	return s.fr.Add(d)
}

func (cmd FileUploadRecordAddCmd) toFileUploadRecord(
	d *domain.FileUploadRecord,
) {
	var createAt int64
	if createAt = cmd.CreateAt; cmd.CreateAt == 0 {
		createAt = getUnixLocalTime()
	}
	*d = domain.FileUploadRecord{
		UserName:   cmd.UserName,
		UploadPath: cmd.UploadPath,
		CreateAt:   createAt,
	}
}

func (cmd FileUploadRecordAddCmd) Validate() error {
	fileUpload := cmd.FileUploadRecord

	b := fileUpload.UserName == nil ||
		fileUpload.UploadPath == "" ||
		fileUpload.CreateAt == 0
	if b {
		return errors.New("invalid cmd of add file upload record")
	}

	return nil
}

type FileUploadRecordDTO struct {
	Users    []string `json:"users"`
	Counts   int64    `json:"counts"`
	UpdateAt string   `json:"update_at"`
}

type FileUploadRecordAddCmd struct {
	domain.FileUploadRecord
}
