package pgsql

import (
	"context"
)

func (cli *client) create(
	ctx context.Context, table interface{},
	data interface{},
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Create(data).Error
}

func (cli *client) filter(
	ctx context.Context, table interface{},
	t string, b *[]BigModelRecord,
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Where("bigmodel=?", t).
		Find(&b).Error
}

func (cli *client) all(
	ctx context.Context, table interface{},
	result interface{},
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Find(&result).Error
}

func (cli *client) distinct(
	ctx context.Context, table interface{},
	d string, result *[]interface{},
) error {
	res := cli.db.WithContext(ctx).
		Model(table).
		Distinct(d).
		Find(&result)
	return res.Error
}
