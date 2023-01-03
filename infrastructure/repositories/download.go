package repositories

import "project/xihe-statistics/domain/repository"

type DownloadRecordMapper interface {
	GetDownloadCount() (int64, error)
}

func NewDownloadRecordRepository(mapper DownloadRecordMapper) repository.DownloadRecord {
	return downloadRecord{mapper}
}

type downloadRecord struct {
	mapper DownloadRecordMapper
}

func (impl downloadRecord) Get() (count int64, err error) {
	return impl.mapper.GetDownloadCount()
}