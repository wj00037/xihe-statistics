package user

import "project/xihe-statistics/domain"

type D1Service interface {
	AddUserWithRepo(*domain.UserWithRepo) error
}