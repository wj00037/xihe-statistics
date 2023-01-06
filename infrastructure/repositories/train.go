package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type TrainRecordMapper interface {
	Get() (int64, error)
	Add(do TrainRecordDO) error
}

func NewTrainRecordRepository(mapper TrainRecordMapper) repository.TrainRecord {
	return trainRecord{mapper}
}

type trainRecord struct {
	mapper TrainRecordMapper
}

func (impl trainRecord) Get() (counts int64, err error) {
	return impl.mapper.Get()
}

func (impl trainRecord) Add(tr *domain.TrainRecord) error {
	return impl.mapper.Add(impl.toTrainRecordDO(tr))
}

type TrainRecordDO struct {
	UserName  string `json:"username"`
	ProjectId string `json:"project_id"`
	TrainId   string `json:"train_id"`
	CreateAt  int64  `json:"create_at"`
}

func (impl trainRecord) toTrainRecordDO(tr *domain.TrainRecord) (do TrainRecordDO) {
	do = TrainRecordDO{
		UserName:  tr.UserName,
		ProjectId: tr.ProjectId,
		TrainId:   tr.TrainId,
		CreateAt:  tr.CreateAt,
	}

	return
}
