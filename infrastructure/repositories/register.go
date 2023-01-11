package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type RegisterRecordMapper interface {
	Add(RegisterRecordDO) error
	Counts() (int64, error)
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

func (impl registerRecord) Get() (do repository.RegisterCounts, err error) {
	counts, err := impl.mapper.Counts()
	if err != nil {
		return
	}
	do = repository.RegisterCounts{
		Counts: counts,
	}
	return
}

func (impl registerRecord) toRegisterRecordDO(d *domain.RegisterRecord) RegisterRecordDO {
	do := RegisterRecordDO{
		UserName: d.UserName.Account(),
		CreateAt: d.CreateAt,
	}
	return do
}
