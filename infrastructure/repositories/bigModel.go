package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type BigModelMapper interface {
	Add(BigModelDO) error
	Get(string) ([]BigModelDO, error)
	GetAll() ([]BigModelDO, error)
}

func NewBigModelRecordRepository(mapper BigModelMapper) repository.UserWithBigModel {
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
	return impl.mapper.Add(impl.toBigmodelRecordDO(d))
}

func (impl bigmodel) Get(d domain.BigModel) (ds []domain.UserWithBigModel, err error) {
	dos, err := impl.mapper.Get(d.BigModel())

	ds = make([]domain.UserWithBigModel, len(dos))
	for j := range dos {
		dos[j].toBigModel(&ds[j])
	}
	if err != nil {
		return
	}

	return
}

func (impl bigmodel) GetAll() (ds []domain.UserWithBigModel, err error) {
	dos, err := impl.mapper.GetAll()

	ds = make([]domain.UserWithBigModel, len(dos))
	for j := range dos {
		dos[j].toBigModel(&ds[j])
	}
	if err != nil {
		return
	}

	return
}

func (impl bigmodel) toBigmodelRecordDO(d *domain.UserWithBigModel) BigModelDO {
	do := BigModelDO{
		UserName: d.UserName,
		BigModel: d.BigModel.BigModel(),
		CreateAt: d.CreateAt,
	}

	return do
}

func (b *BigModelDO) toBigModel(d *domain.UserWithBigModel) (err error) {
	d.UserName = b.UserName
	d.BigModel, err = domain.NewBigModel(b.BigModel)
	d.CreateAt = b.CreateAt
	return
}
