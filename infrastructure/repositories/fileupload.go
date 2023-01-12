package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type FileUploadRecordMapper interface {
	GetUsers() (FileUploadUserCountsDO, error)
	AddRecord(FileUploadRecordDO) error
}

func NewFileUploadRecordRepository(mapper FileUploadRecordMapper) repository.FileUploadRecord {
	return fileUploadRecord{mapper}
}

type fileUploadRecord struct {
	mapper FileUploadRecordMapper
}

func (impl fileUploadRecord) Get() (f repository.FileUploadUsers, err error) {
	do, err := impl.mapper.GetUsers()
	if err != nil {
		return
	}

	f = repository.FileUploadUsers{
		Users: do.Users,
	}

	return
}

func (impl fileUploadRecord) Add(d *domain.FileUploadRecord) (err error) {
	return impl.mapper.AddRecord(impl.toFileUploadRecordDO(d))
}

func (impl fileUploadRecord) toFileUploadRecordDO(
	d *domain.FileUploadRecord,
) FileUploadRecordDO {
	return FileUploadRecordDO{
		UserName:   d.UserName.Account(),
		UploadPath: d.UploadPath,
		CreateAt:   d.CreateAt,
	}
}

type FileUploadUserCountsDO struct {
	Users []string
}

type FileUploadRecordDO struct {
	UserName   string
	UploadPath string
	CreateAt   int64
}
