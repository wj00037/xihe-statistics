package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type CloudRecordMapper interface {
	AddCloudRecord(*CloudRecordDO) error
}

func NewCloudRecordRepository(mapper CloudRecordMapper) repository.CloudRecord {
	return &cloudRecord{mapper}
}

type cloudRecord struct {
	mapper CloudRecordMapper
}

func (impl *cloudRecord) Add(d *domain.Cloud) (err error) {
	do := new(CloudRecordDO)
	do.toCloudRecordDO(d)

	return impl.mapper.AddCloudRecord(do)
}

func (do *CloudRecordDO) toCloudRecordDO(d *domain.Cloud) {
	*do = CloudRecordDO{
		UserName: d.UserName.Account(),
		CloudId:  d.CloudId,
		CreateAt: d.CreateAt,
	}
}

type CloudRecordDO struct {
	UserName string
	CloudId  string
	CreateAt int64
}
