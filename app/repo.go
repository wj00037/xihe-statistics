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
	Get() (RepoRecordDTO, error)
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

func (s repoRecordService) Get() (dto RepoRecordDTO, err error) {
	rr, err := s.ur.Get()
	if err != nil {
		return
	}

	dto = RepoRecordDTO{
		Users:    rr.Users,
		Counts:   rr.Counts,
		UpdateAt: time.Now().Format("2006-01-02 15:04:05+08:00"),
	}
	return
}

func (cmd RepoRecordAddCmd) toRepo(ur *domain.UserWithRepo) {
	now := time.Now().Unix()

	*ur = domain.UserWithRepo{
		UserName: cmd.UserName,
		RepoName: cmd.RepoName,
		CreateAt: now,
	}
}

type RepoRecordDTO struct {
	Users    []string `json:"users"`
	Counts   int      `json:"counts"`
	UpdateAt string   `json:"update_at"`
}
