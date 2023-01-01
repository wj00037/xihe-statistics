package pgsql

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"
)

func NewFileUploadRecordMapper(table FileUploadRecord) repositories.FileUploadRecordMapper {
	return fileUploadRecord{table}
}

type fileUploadRecord struct {
	table FileUploadRecord
}

func (m fileUploadRecord) GetUsers() (
	do repositories.FileUploadUserCountsDO,
	err error,
) {

	var users []interface{}

	f := func(ctx context.Context) error {
		return cli.distinct(
			ctx, m.table,
			"username", &users,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	do = repositories.FileUploadUserCountsDO{
		Users: toArryString(users),
	}

	return
}

func (m fileUploadRecord) AddRecord(
	do repositories.FileUploadRecordDO,
) (err error) {

	data, _ := m.toFileUploadCol(do)

	f := func(ctx context.Context) error {
		return cli.create(
			ctx, m.table,
			data,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (m fileUploadRecord) toFileUploadCol(
	do repositories.FileUploadRecordDO,
) (FileUploadRecord, error) {
	colObj := FileUploadRecord{
		UserName:   do.UserName,
		UploadPath: do.UploadPath,
		CreateAt:   do.CreateAt,
	}

	return colObj, nil
}
