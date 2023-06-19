package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type BigModelMapper interface {
	Add(BigModelDO) error
	Get(string) ([]BigModelDO, error)
	GetByTypeAndTime(string, int64) (int64, error)
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
	if err != nil {
		return
	}

	ds = make([]domain.UserWithBigModel, len(dos))
	for j := range dos {
		if err = dos[j].toBigModel(&ds[j]); err != nil {
			return
		}
	}

	return
}

func (impl bigmodel) GetByTypeAndTime(d domain.BigModel, time int64) (int64, error) {
	return impl.mapper.GetByTypeAndTime(d.BigModel(), time)
}

func (impl bigmodel) GetAll() (ds []domain.UserWithBigModel, err error) {
	dos, err := impl.mapper.GetAll()
	if err != nil {
		return
	}

	ds = make([]domain.UserWithBigModel, len(dos))
	for j := range dos {
		if err = dos[j].toBigModel(&ds[j]); err != nil {
			return
		}
	}

	return
}

func (impl bigmodel) toBigmodelRecordDO(d *domain.UserWithBigModel) BigModelDO {
	do := BigModelDO{
		UserName: d.UserName.Account(),

		BigModel: d.BigModel.BigModel(),
		CreateAt: d.CreateAt,
	}

	return do
}

func (b *BigModelDO) toBigModel(d *domain.UserWithBigModel) (err error) {
	if d.UserName, err = domain.NewAccount(b.UserName); err != nil {
		return
	}
	d.BigModel, err = domain.NewBigModel(b.BigModel)
	d.CreateAt = b.CreateAt

	return
}
