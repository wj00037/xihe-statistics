package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

// BigModel
type BigModelRecordMessageService interface {
	AddUserWithBigModel(*UserWithBigModelAddCmd) error
}

func NewBigModelRecordMessageService(
	ub repository.UserWithBigModel,
) BigModelRecordMessageService {
	return bigModelRecordMessageService{
		ub: ub,
	}
}

type bigModelRecordMessageService struct {
	ub repository.UserWithBigModel
}

func (b bigModelRecordMessageService) AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error { // implement app function with infrastructure function
	v := new(domain.UserWithBigModel)
	cmd.toBigModel(v)

	err := b.ub.Add(v)
	if err != nil {
		return err
	}
	return nil
}

// Repo
type RepoRecordMessageService interface {
	Add(*RepoRecordAddCmd) error
}

func NewRepoRecordMessageService(
	ur repository.UserWithRepo,
) RepoRecordMessageService {
	return repoRecordMessageService{
		ur: ur,
	}
}

type repoRecordMessageService struct {
	ur repository.UserWithRepo
}

func (s repoRecordMessageService) Add(cmd *RepoRecordAddCmd) (err error) {
	uwr := new(domain.UserWithRepo)
	cmd.toRepo(uwr)

	return s.ur.Add(uwr)
}
