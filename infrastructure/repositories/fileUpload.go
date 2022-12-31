package repositories

import "project/xihe-statistics/domain/repository"

type FileUploadRecordMapper interface {
	GetUsers() (FileUploadUserCountsDO, error)
}

func NewFileUploadRecordRepository(mapper FileUploadRecordMapper) repository.FileUploadRecord {
	return fileUploadRecord{mapper}
}

type fileUploadRecord struct {
	mapper FileUploadRecordMapper
}

type FileUploadUserCountsDO struct {
	Users []string
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
