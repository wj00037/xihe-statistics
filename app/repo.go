package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
	"time"
)

type RepoRecordAddCmd struct {
	domain.UserWithRepo
}

type RepoRecordService interface {
	Add(*RepoRecordAddCmd) error
}

func NewRepoRecordService(
	ur repository.UserWithRepo,
) RepoRecordService {
	return repoRecordService{
		ur: ur,
	}
}

type repoRecordService struct {
	ur repository.UserWithRepo
}

func (s repoRecordService) Add(cmd *RepoRecordAddCmd) (err error) {
	uwr := new(domain.UserWithRepo)
	cmd.toRepo(uwr)

	return s.ur.Add(uwr)
}

func (cmd RepoRecordAddCmd) toRepo(ur *domain.UserWithRepo) {
	now := time.Now().Unix()

	*ur = domain.UserWithRepo{
		UserName: cmd.UserName,
		RepoName: cmd.RepoName,
		CreateAt: now,
	}

}
