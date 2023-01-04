package pgsql

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"
)

func NewTrainRecordMapper(table TrainRecord) repositories.TrainRecordMapper {
	return trainRecord{table}
}

type trainRecord struct {
	table TrainRecord
}

func (m trainRecord) Get() (counts int64, err error) {

	f := func(ctx context.Context) error {
		return cli.count(
			ctx, m.table,
			&counts,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}
