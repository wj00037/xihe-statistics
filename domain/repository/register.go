package repository

import "project/xihe-statistics/domain"

type RegisterRecord interface {
	Add(*domain.RegisterRecord) error
}