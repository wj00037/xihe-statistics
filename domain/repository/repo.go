package repository

import (
	"project/xihe-statistics/domain"
)

type RepoRecords struct {
	Users  []string
	Counts int
}

type UserWithRepo interface {
	Add(*domain.UserWithRepo) error
	Get() (RepoRecords, error)
}
