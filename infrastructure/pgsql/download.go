package pgsql

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"
)

func NewDownloadRecordMapper(table DownloadRecord) repositories.DownloadRecordMapper {
	return downloadRecord{table}
}

type downloadRecord struct {
	table DownloadRecord
}

func (m downloadRecord) GetDownloadCount() (count int64, err error) {
	f := func(ctx context.Context) error {
		return cli.count(
			ctx, m.table,
			&count,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (m downloadRecord) AddDownloadRecord(do repositories.DownloadRecordDO) (err error) {

	col := m.toDownloadRecordCol(do)

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

func (m downloadRecord) toDownloadRecordCol(
	do repositories.DownloadRecordDO,
) (
	col DownloadRecord,
) {
	col = DownloadRecord{
		UserName:     do.UserName,
		DownloadPath: do.DownloadPath,
		CreateAt:     do.CreateAt,
	}

	return
}
