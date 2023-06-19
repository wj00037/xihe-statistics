package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type DownloadRecordMapper interface {
	AddDownloadRecord(DownloadRecordDO) error
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

func (impl downloadRecord) Add(dr *domain.DownloadRecord) (err error) {
	do := impl.toDownloadRecordDO(dr)

	return impl.mapper.AddDownloadRecord(do)
}

type DownloadRecordDO struct {
	UserName     string
	DownloadPath string
	CreateAt     int64
}

func (impl downloadRecord) toDownloadRecordDO(dr *domain.DownloadRecord) DownloadRecordDO {
	return DownloadRecordDO{
		UserName:     dr.UserName.Account(),
		DownloadPath: dr.DownloadPath,
		CreateAt:     dr.CreateAt,
	}
}
