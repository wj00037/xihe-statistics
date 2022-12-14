package repository

import (
	"project/xihe-statistics/domain"
)

type UserWithBigModel interface {
	Add(*domain.UserWithBigModel) error
	Get(domain.BigModel) ([]domain.UserWithBigModel, error)
	GetAll() ([]domain.UserWithBigModel, error)
}
