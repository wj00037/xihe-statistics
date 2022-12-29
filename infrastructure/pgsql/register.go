package pgsql

import (
	"context"

	"project/xihe-statistics/infrastructure/repositories"
)

func NewRegisterRecordMapper(table RegisterRecord) repositories.RegisterRecordMapper {
	return registerRecord{table}
}

type registerRecord struct {
	table RegisterRecord
}

func (m registerRecord) Add(
	b repositories.RegisterRecordDO,
) (err error) {
	col, err := m.toRegisterRecordCol(b)
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

func (m registerRecord) Counts() (
	count int64, err error,
) {
	f := func(ctx context.Context) error {
		return cli.count(
			ctx, m.table, &count,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (m registerRecord) toRegisterRecordCol(b repositories.RegisterRecordDO) (RegisterRecord, error) {
	ColObj := RegisterRecord{
		UserName: b.UserName,
		CreateAt: b.CreateAt,
	}
	return ColObj, nil
}
