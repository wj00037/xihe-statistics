package app

import "project/xihe-statistics/domain/repository"

type FileUploadRecordService interface {
	GetUsersCounts() (FileUploadRecordDTO, error)
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

type FileUploadRecordDTO struct {
	Users    []string `json:"users"`
	Counts   int64    `json:"counts"`
	UpdateAt string   `json:"update_at"`
}
