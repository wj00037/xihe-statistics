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

func (m trainRecord) GetTrains(startTime int64, endTime int64) (counts int64, err error) {

	f := func(ctx context.Context) error {
		return cli.whereCount(
			ctx, m.table,
			"create_at >= ? AND create_at <= ?",
			startTime, endTime,
			&counts,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (m trainRecord) Add(do repositories.TrainRecordDO) (err error) {

	col := m.toTrainRecordCol(do)

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

func (m trainRecord) toTrainRecordCol(
	do repositories.TrainRecordDO,
) TrainRecord {
	return TrainRecord{
		UserName:  do.UserName,
		ProjectId: do.ProjectId,
		TrainId:   do.TrainId,
		CreateAt:  do.CreateAt,
	}
}
