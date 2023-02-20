package app

import (
	"errors"
	"project/xihe-statistics/domain/platform"
	"project/xihe-statistics/domain/repository"
	"time"
)

type GitLabService interface {
	Counts() (CloneCountsDTO, error)
	Save(*CloneCountsCmd) error
}

type gitLabService struct {
	pf platform.PlatForm
	gl repository.Gitlab
}

func NewGitLabService(
	pf platform.PlatForm,
	gl repository.Gitlab,
) GitLabService {
	return &gitLabService{
		pf: pf,
		gl: gl,
	}
}

func (g *gitLabService) Counts() (dto CloneCountsDTO, err error) {
	var (
		c      = 1
		counts int64
	)
	for {
		count, err := g.countsPage(c)
		if err != nil {
			if IsErrorEmptyProjectIdPage(err) {
				return CloneCountsDTO{
					Counts:   counts,
					CreateAt: getUnixLocalTime(),
				}, nil
			}

			return CloneCountsDTO{}, err
		}

		counts += count
		c++
		time.Sleep(time.Second)
	}
}

func (g *gitLabService) countsPage(pageNum int) (counts int64, err error) {
	ids, err := g.pf.GetProjectId(pageNum)
	if err != nil {
		return
	}

	if len(ids) == 0 {
		err = errorEmptyGitLabProjectIdPage{
			errors.New(ErrorEmptyGitLabProjectIdPage),
		}

		return
	}

	for _, id := range ids {
		total, err := g.pf.GetCloneTotal(id.Id)
		if err != nil {
			return 0, err
		}

		counts += total.Total
	}

	return
}

func (g *gitLabService) Save(cmd *CloneCountsCmd) error {
	cc, err := g.toCloneCount(cmd)
	if err != nil {
		return err
	}
	return g.gl.InsertCloneCount(&cc)
}

type CloneCountsDTO struct {
	Counts   int64 `json:"counts"`
	CreateAt int64 `json:"create_at"`
}

type CloneCountsCmd struct {
	Counts   int64 `json:"counts"`
	CreateAt int64 `json:"create_at"`
}

func (g *gitLabService) toCloneCount(cmd *CloneCountsCmd) (repository.CloneCount, error) {
	return repository.CloneCount{
		Counts:   cmd.Counts,
		CreateAt: cmd.CreateAt,
	}, nil
}
