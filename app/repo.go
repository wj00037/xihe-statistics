package app

import (
	"errors"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
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
		UpdateAt: getLocalTime(),
	}
	return
}

func (cmd RepoRecordAddCmd) Validate() error {
	repo := cmd.UserWithRepo

	b := repo.UserName == nil ||
		repo.RepoName == "" ||
		repo.CreateAt == 0

	if b {
		return errors.New("invalid cmd of add repo record")
	}

	return nil
}

func (cmd RepoRecordAddCmd) toRepo(ur *domain.UserWithRepo) {
	var createAt int64
	if createAt = cmd.CreateAt; cmd.CreateAt == 0 {
		createAt = GetUnixLocalTime()
	}

	*ur = domain.UserWithRepo{
		UserName: cmd.UserName,
		RepoName: cmd.RepoName,
		CreateAt: createAt,
	}
}

type RepoRecordDTO struct {
	Users    []string `json:"users"`
	Counts   int      `json:"counts"`
	UpdateAt string   `json:"update_at"`
}
