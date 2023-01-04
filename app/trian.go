package app

import (
	"project/xihe-statistics/domain/repository"
)

type TrainRecordService interface {
	Get() (TrainRecordDTO, error)
}

type trainRecordService struct {
	tr repository.TrainRecord
}

func NewTrainRecordService(
	tr repository.TrainRecord,
) TrainRecordService {
	return trainRecordService{
		tr: tr,
	}
}

func (s trainRecordService) Get() (
	dto TrainRecordDTO,
	err error,
) {
	counts, err := s.tr.Get()

	if err != nil {
		return
	}

	dto = TrainRecordDTO{
		Counts:   counts,
		UpdateAt: getLocalTime(),
	}

	return
}

type TrainRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}
