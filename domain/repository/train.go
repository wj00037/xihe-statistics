package repository

import "project/xihe-statistics/domain"

type TrainRecord interface {
	Add(*domain.TrainRecord) error
	Get() (int64, error)
}
