package message

import "project/xihe-statistics/domain"

type BigModelRecordHandler interface {
	AddBigModelRecord(*domain.UserWithBigModel) error
}

type RepoRecordHandler interface {
	AddRepoRecord(*domain.UserWithRepo) error
}