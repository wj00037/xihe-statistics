package app

import (
	"errors"
	"fmt"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/user"
	"time"
)

func (cmd *UserWithBigModelAddCmd) Validate() error {
	err := errors.New("invalid cmd of add user query big model record")

	b := cmd.UserName != ""

	if !b {
		return err
	}

	_, err = domain.NewBigModel(cmd.BigModel.BigModel())

	if err != nil {
		return err
	}

	return nil
}

type UserWithBigModelAddCmd struct {
	domain.UserWithBigModel
}

type BigModelRecordService interface {
	AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error
}

func NewBigModelRecordService(
	ub user.UserWithBigModel,
) BigModelRecordService {
	return bigModelRecordService{
		ub: ub,
	}
}

type bigModelRecordService struct {
	ub user.UserWithBigModel
}

func (b bigModelRecordService) AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error {
	v := new(domain.UserWithBigModel)
	cmd.toBigModel(v)

	fmt.Printf("v: %v\n", v)

	err := b.ub.Add(v)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *UserWithBigModelAddCmd) toBigModel(r *domain.UserWithBigModel) {
	now := time.Now().Unix()

	*r = domain.UserWithBigModel{
		UserName: cmd.UserName,
		BigModel: cmd.BigModel,
		CreateAt: now,
	}
}
