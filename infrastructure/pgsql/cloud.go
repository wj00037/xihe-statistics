package pgsql

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"
)

func NewCloudRecordMapper(table CloudRecord) repositories.CloudRecordMapper {
	return &cloudRecord{table}
}

type cloudRecord struct {
	table CloudRecord
}

func (m *cloudRecord) AddCloudRecord(do *repositories.CloudRecordDO) (err error) {
	col := toCloudRecordCol(do)

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

func (m *cloudRecord) GetCloudRecordCount() (counts int64, err error) {
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

func toCloudRecordCol(do *repositories.CloudRecordDO) CloudRecord {

	return CloudRecord{
		UserName: do.UserName,
		CloudId:  do.CloudId,
		CreateAt: do.CreateAt,
	}
}
