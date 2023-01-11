package app

import (
	"errors"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type RegisterRecordService interface {
	Add(*RegisterRecordAddCmd) error
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

func (s registerRecordService) Add(cmd *RegisterRecordAddCmd) (err error) {
	register := new(domain.RegisterRecord)
	cmd.toRegisterRecord(register)

	return s.rr.Add(register)
}

func (s registerRecordService) Get() (
	dto RegisterRecordDTO, err error,
) {
	rc, err := s.rr.Get()
	dto = RegisterRecordDTO{
		Counts:   rc.Counts,
		UpdateAt: getLocalTime(),
	}

	return
}

func (cmd RegisterRecordAddCmd) toRegisterRecord(
	d *domain.RegisterRecord,
) {
	register := cmd.RegisterRecord

	*d = domain.RegisterRecord{
		UserName: register.UserName,
		CreateAt: register.CreateAt,
	}
}

func (cmd RegisterRecordAddCmd) Validate() error {
	register := cmd.RegisterRecord
	b := register.UserName == nil ||
		register.CreateAt == 0

	if b {
		return errors.New("invalid cmd of add register record")
	}

	return nil
}

type RegisterRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}

type RegisterRecordAddCmd struct {
	domain.RegisterRecord
}
