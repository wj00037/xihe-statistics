package app

import (
	"errors"

	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type TrainRecordService interface {
	Add(*TrainRecordAddCmd) error
	Get() (TrainRecordDTO, error)
	GetTrains(TrainIncreaseCmd) (dto TrainsDTO, err error)
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

func (s trainRecordService) Add(cmd *TrainRecordAddCmd) (err error) {
	tr := new(domain.TrainRecord)
	cmd.toTrainRecord(tr)

	return s.tr.Add(tr)
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

func (s trainRecordService) GetTrains(
	cmd TrainIncreaseCmd,
) (dto TrainsDTO, err error) {
	startTime := cmd.StartTime
	endTime := cmd.EndTime

	tStart, err := toUnixTime(startTime)
	if err != nil {
		return
	}

	tEnd, err := toUnixTime(endTime)
	if err != nil {
		return
	}

	counts, err := s.tr.GetTrains(tStart, tEnd)
	if err != nil {
		return
	}

	dto = TrainsDTO{
		Counts: counts,
	}

	return
}

func (cmd TrainRecordAddCmd) toTrainRecord(
	d *domain.TrainRecord,
) {
	*d = domain.TrainRecord{
		UserName:  cmd.UserName,
		ProjectId: cmd.ProjectId,
		TrainId:   cmd.TrainId,
		CreateAt:  cmd.CreateAt,
	}
}

func (cmd TrainRecordAddCmd) Validate() error {
	b := cmd.UserName == "" ||
		cmd.ProjectId == "" ||
		cmd.TrainId == "" ||
		cmd.CreateAt == 0

	if b {
		return errors.New("invalid cmd of add train record")
	}

	return nil
}

type TrainRecordAddCmd struct {
	domain.TrainRecord
}

type TrainIncreaseCmd struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type TrainsDTO struct {
	Counts int64 `json:"counts"`
}

type TrainRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}
