package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type RegisterRecordMapper interface {
	Add(RegisterRecordDO) error
}

type RegisterRecordDO struct {
	UserName string
	CreateAt int64
}

type registerRecord struct {
	mapper RegisterRecordMapper
}

func NewRegisterRecordRepository(mapper RegisterRecordMapper) repository.RegisterRecord {
	return registerRecord{mapper}
}

func (impl registerRecord) Add(d *domain.RegisterRecord) (err error) {
	return impl.mapper.Add(impl.toRegisterRecordDO(d))
}

func (impl registerRecord) toRegisterRecordDO(d *domain.RegisterRecord) RegisterRecordDO {
	do := RegisterRecordDO{
		UserName: d.UserName,
		CreateAt: d.CreateAt,
	}
	return do
}
