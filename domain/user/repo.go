package user

import "project/xihe-statistics/domain"

type UserWithRepo interface {
	Add(*domain.UserWithRepo) error
}