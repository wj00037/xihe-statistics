package pgsql

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"
)

func NewMediaMapper(table Media) repositories.MediaMapper {
	return mediaModel{table}
}

type mediaModel struct {
	table Media
}

func (m mediaModel) Get(t string) (do repositories.MediaDataDO, err error) {

	f := func(ctx context.Context) error {
		return cli.whereCount2(
			ctx, m.table, "name = ?", t,
			&do.Counts)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (m mediaModel) Add(name string, time int64) (err error) {

	media := Media{
		Name:     name,
		CreateAt: time,
	}

	f := func(ctx context.Context) error {
		return cli.create(
			ctx, m.table, media,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}
