package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type MediaService interface {
	Add(*AddMediaCmd) error
	GetAll() (AllMediaDataDTO, error)
}

func NewMeidaService(
	repo repository.Media,
) MediaService {
	return &mediaService{
		repo: repo,
	}
}

type mediaService struct {
	repo repository.Media
}

func (s *mediaService) Add(cmd *AddMediaCmd) error {

	m := &domain.Media{
		Name:     cmd.Name,
		CreateAt: cmd.CreateAt,
	}

	return s.repo.Add(m)
}

func (s *mediaService) GetAll() (dto AllMediaDataDTO, err error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return
	}

	dto = toAllMediaDataDTO(&data)

	return
}
