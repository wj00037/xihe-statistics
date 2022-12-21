package pgsql

import "context"

func (cli *client) create(
	ctx context.Context, table interface{},
	data interface{},
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Save(data).Error
}
