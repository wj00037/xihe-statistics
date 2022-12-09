package user

import "project/xihe-statistics/domain"

type UserWithBigModel interface {
	Add(*domain.UserWithBigModel) error
}
