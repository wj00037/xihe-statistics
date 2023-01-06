package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type DownloadRecordService interface {
	Add(*domain.DownloadRecord) error
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

func (s downloadRecordService) Add(d *domain.DownloadRecord) (err error) {
	return s.dr.Add(d)
}

func (s downloadRecordService) Get() (dto DownloadRecordDTO, err error) {
	counts, err := s.dr.Get()
	dto = DownloadRecordDTO{
		Counts:   counts,
		UpdateAt: getLocalTime(),
	}
	return
}

type DownloadRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}
