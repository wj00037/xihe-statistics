package app

import (
	"time"

	"project/xihe-statistics/domain/repository"
)

type RegisterRecordService interface {
	Get() (RegisterRecordDTO, error)
}

func NewRegisterRecordService(
	rr repository.RegisterRecord,
) RegisterRecordService {
	return registerRecordService{
		rr: rr,
	}
}

type registerRecordService struct {
	rr repository.RegisterRecord
}

func (s registerRecordService) Get() (
	dto RegisterRecordDTO, err error,
) {
	rc, err := s.rr.Get()
	dto = RegisterRecordDTO{
		Counts:   rc.Counts,
		UpdateAt: time.Now().Format("2006-01-02 15:04:05+08:00"),
	}

	return
}

type RegisterRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}
