package pgsql

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"
)

func NewBigModelMapper(table BigModelRecord) repositories.BigModelMapper {
	return bigModel{table}
}

type bigModel struct {
	table BigModelRecord
}

func (m bigModel) Add(
	b repositories.BigModelDO,
) (err error) {
	col, err := m.toBigModelCol(b)
	if err != nil {
		return
	}

	f := func(ctx context.Context) error {
		return cli.create(
			ctx, m.table,
			col,
		)
	}

	if err = withContext(f); err != nil {
		return err
	}

	return
}

func (m bigModel) Get(string) (do []repositories.BigModelDO, err error) {
	return
}

func (m bigModel) GetAll() (do []repositories.BigModelDO, err error) {
	return
}

func (m bigModel) toBigModelCol(b repositories.BigModelDO) (BigModelRecord, error) {
	ColObj := BigModelRecord{
		UserName: b.UserName,
		BigModel: b.BigModel,
		CreateAt: b.CreateAt,
	}

	return ColObj, nil
}

func (m bigModel) toBigModelDO(v BigModelRecordItem, do *repositories.BigModelDO) {
	*do = repositories.BigModelDO{
		UserName: v.UserName,
		BigModel: v.BigModel,
		CreateAt: v.CreateAt,
	}
}
