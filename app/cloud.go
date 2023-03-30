package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type CloudRecordService interface {
	Add(*CloudRecordCmd) error
	Get() (CloudRecordDTO, error)
}

func NewCloudRecodeService(
	repo repository.CloudRecord,
) CloudRecordService {
	return &cloudRecordService{
		repo: repo,
	}
}

type cloudRecordService struct {
	repo repository.CloudRecord
}

func (s *cloudRecordService) Add(cmd *CloudRecordCmd) error {

	return s.repo.Add(
		&domain.Cloud{
			UserName: cmd.User,
			CloudId:  cmd.CloudId,
			CreateAt: cmd.CreateAt,
		},
	)
}

func (s *cloudRecordService) Get() (dto CloudRecordDTO, err error) {
	c, err := s.repo.Get()
	if err != nil {
		return
	}

	dto.toCloudRecordDTO(c, getLocalTime())

	return
}
