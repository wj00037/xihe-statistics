package app

import (
	"errors"

	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type DownloadRecordService interface {
	Add(*DownloadRecordAddCmd) error
	Get() (DownloadRecordDTO, error)
}

type downloadRecordService struct {
	dr repository.DownloadRecord
}

func NewDownloadRecordService(
	dr repository.DownloadRecord,
) DownloadRecordService {
	return downloadRecordService{
		dr: dr,
	}
}

func (s downloadRecordService) Add(cmd *DownloadRecordAddCmd) (err error) {
	download := new(domain.DownloadRecord)
	cmd.toDownloadRecord(download)

	return s.dr.Add(download)
}

func (s downloadRecordService) Get() (dto DownloadRecordDTO, err error) {
	counts, err := s.dr.Get()
	dto = DownloadRecordDTO{
		Counts:   counts,
		UpdateAt: getLocalTime(),
	}
	return
}

func (cmd *DownloadRecordAddCmd) toDownloadRecord(
	d *domain.DownloadRecord,
) {
	*d = domain.DownloadRecord{
		UserName:     cmd.UserName,
		DownloadPath: cmd.DownloadPath,
		CreateAt:     cmd.CreateAt,
	}
}

func (cmd DownloadRecordAddCmd) Validate() error {
	b := cmd.UserName == nil ||
		cmd.DownloadPath == "" ||
		cmd.CreateAt == 0

	if b {
		return errors.New("invalid cmd of add download record")
	}

	return nil
}

type DownloadRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}

type DownloadRecordAddCmd struct {
	domain.DownloadRecord
}
