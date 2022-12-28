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
		return
	}

	return
}

func (m bigModel) Get(t string) (cols []repositories.BigModelDO, err error) {

	var records []BigModelRecord

	f := func(ctx context.Context) error {
		return cli.filter(
			ctx, m.table,
			t, &records,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	cols = make([]repositories.BigModelDO, len(records))
	for j := range records {
		m.toBigModelDO(records[j], &cols[j])
	}

	return
}

func (m bigModel) GetAll() (cols []repositories.BigModelDO, err error) {
	var records []BigModelRecord

	f := func(ctx context.Context) error {
		return cli.all(
			ctx, m.table,
			&records,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	cols = make([]repositories.BigModelDO, len(records))
	for j := range records {
		m.toBigModelDO(records[j], &cols[j])
	}

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

func (m bigModel) toBigModelDO(v BigModelRecord, do *repositories.BigModelDO) {
	*do = repositories.BigModelDO{
		UserName: v.UserName,
		BigModel: v.BigModel,
		CreateAt: v.CreateAt,
	}
}
