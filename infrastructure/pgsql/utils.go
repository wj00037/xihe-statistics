package pgsql

import (
	"context"

	"gorm.io/gorm/clause"
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
	condition string, t string,
	b *[]BigModelRecord,
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Where(condition, t).
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

func (cli *client) count(
	ctx context.Context, table interface{},
	counts *int64,
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Count(counts).Error
}

func (cli *client) distinctCount(
	ctx context.Context, table interface{},
	d string, counts *int64,
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Distinct(d).
		Count(counts).Error
}

func (cli *client) fileUploadUpsert(
	ctx context.Context, table interface{},
	data FileUploadRecord,
) error {
	return cli.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		UpdateAll: true,
	}).Create(&data).Error
}

// filter data with where command (two condtion) and count distinct data return quantity
func (cli *client) whereDistinctCount(
	ctx context.Context, table interface{},
	condition string, v1 interface{}, v2 interface{},
	d string, counts *int64,
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Where(condition, v1, v2).
		Distinct(d).
		Count(counts).Error
}

func (cli *client) whereCount(
	ctx context.Context, table interface{},
	condition string, v1 interface{}, v2 interface{},
	counts *int64,
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Where(condition, v1, v2).
		Count(counts).Error
}

func (cli *client) getLast(
	ctx context.Context, table interface{},
	result interface{},
) error {
	return cli.db.WithContext(ctx).
		Model(table).
		Last(result).Error
}
