package app

import (
	"errors"
	"project/xihe-statistics/domain"
)

type CloudRecordCmd struct {
	User     domain.Account
	CloudId  string
	CreateAt int64
}

func (cmd CloudRecordCmd) Validate() error {
	b := cmd.User == nil ||
		cmd.CloudId == "" ||
		cmd.CreateAt == 0

	if b {
		return errors.New("invalid cmd of add cloud record")
	}

	return nil
}

type CloudRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}

func (dto *CloudRecordDTO) toCloudRecordDTO(counts int64, update string) {
	*dto = CloudRecordDTO{
		Counts:   counts,
		UpdateAt: update,
	}
}
