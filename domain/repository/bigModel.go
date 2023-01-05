package repository

import (
	"project/xihe-statistics/domain"
)

type UserWithBigModel interface {
	Add(*domain.UserWithBigModel) error
	Get(domain.BigModel) ([]domain.UserWithBigModel, error)
	GetByTypeAndTime(domain.BigModel, int64) (int64, error)
	GetAll() ([]domain.UserWithBigModel, error)
}
