package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/user"
)

type BigModelMapper interface {
	Add(BigModelDO) error
}

func NewBigModelRecordRepository(mapper BigModelMapper) user.UserWithBigModel {
	return bigmodel{mapper}
}

type bigmodel struct {
	mapper BigModelMapper
}

type BigModelDO struct {
	UserName string
	BigModel string
	CreateAt int64
}

func (impl bigmodel) Add(d *domain.UserWithBigModel) error {
	err := impl.mapper.Add(impl.toBigmodelRecordDO(d))
	if err != nil {
		return err
	}
	return nil
}

func (impl bigmodel) toBigmodelRecordDO(d *domain.UserWithBigModel) BigModelDO {
	do := BigModelDO{
		UserName: d.UserName,
		BigModel: d.BigModel.BigModel(),
		CreateAt: d.CreateAt,
	}

	return do
}
