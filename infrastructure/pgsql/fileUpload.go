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
