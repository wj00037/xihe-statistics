package repository

import "project/xihe-statistics/domain"

type TrainRecord interface {
	Add(*domain.TrainRecord) error
	GetTrains(startTime int64, endTime int64) (counts int64, err error)
	Get() (int64, error)
}
