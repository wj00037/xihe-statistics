package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

// repo record
type UserWithRepoMapper interface {
	Add(UserWithRepoDO) error
}

func NewUserWithRepoRepository(mapper UserWithRepoMapper) repository.UserWithRepo {
	return userWithRepo{mapper}
}

type userWithRepo struct {
	mapper UserWithRepoMapper
}

type UserWithRepoDO struct {
	UserName string
	RepoName string
	CreateAt int64
}

func (impl userWithRepo) Add(u *domain.UserWithRepo) error {
	return impl.mapper.Add(impl.toUserWithRepoDO(u))
}

func (impl userWithRepo) toUserWithRepoDO(u *domain.UserWithRepo) UserWithRepoDO {
	do := UserWithRepoDO{
		UserName: u.UserName,
		RepoName: u.RepoName,
		CreateAt: u.CreateAt,
	}

	return do
}
