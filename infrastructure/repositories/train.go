package repositories

import "project/xihe-statistics/domain/repository"

type TrainRecordMapper interface {
	Get() (int64, error)
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
