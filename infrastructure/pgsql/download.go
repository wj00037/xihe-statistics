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
