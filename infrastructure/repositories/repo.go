package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

// repo record
type UserWithRepoMapper interface {
	Add(UserWithRepoDO) error
	Get() (RepoRecordsDO, error)
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

type RepoRecordsDO struct {
	Users  []string
	Counts int
}

func (impl userWithRepo) Add(u *domain.UserWithRepo) error {
	return impl.mapper.Add(impl.toUserWithRepoDO(u))
}

func (impl userWithRepo) Get() (r repository.RepoRecords, err error) {
	do, err := impl.mapper.Get()
	r = impl.toRepoRecord(do)
	return
}

func (impl userWithRepo) toUserWithRepoDO(u *domain.UserWithRepo) UserWithRepoDO {
	do := UserWithRepoDO{
		UserName: u.UserName.Account(),
		RepoName: u.RepoName,
		CreateAt: u.CreateAt,
	}

	return do
}

func (impl userWithRepo) toRepoRecord(r RepoRecordsDO) repository.RepoRecords {
	return repository.RepoRecords{
		Users:  r.Users,
		Counts: r.Counts,
	}
}
