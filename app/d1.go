package app

import (
	"errors"

	"github.com/sirupsen/logrus"

	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/user"
)

func (cmd *UserWithRepoCmd) Validate() error {
	err := errors.New("invalid cmd of add user create repo record")

	b := cmd.UserName != ""

	if !b {
		return err
	}

	return nil

}

func (cmd *UserWithBigModelAddCmd) Validate() error {
	err := errors.New("invalid cmd of add user query big model record")

	b := cmd.UserName != "" &&
		cmd.BigModelType != ""

	if !b {
		return err
	}

	t := cmd.BigModelType == "taichu-VQA" ||
		cmd.BigModelType == "taichu-TextToImg" ||
		cmd.BigModelType == "taichu-ImgToText"
	if !t {
		return err
	}

	return nil
}

type D1Service interface {
	AddUserWithRepo(cmd *UserWithRepoCmd) error
	AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error
	GetUserWithRepo() (UserWithRepoDTO, error)
	GetUserWithBigModel() (UserWithBigModelDTO, error)
	GetUsers() (UsersDTO, error)
}

type UserWithRepoCmd struct {
	UserName string `json:"username"`
}

type UserWithBigModelAddCmd struct {
	BigModelType string `json:"bigmodel_type"`
	UserName     string `json:"username"`
}

type UserWithRepoDTO struct {
	Counts   int   `json:"counts"`
	UpdateAt int64 `json:"update_at"`
}

type UserWithBigModelDTO struct {
	BigModelType string        `json:"bigmodel_type"`
	Users        []domain.User `json:"users"`
	Counts       int           `json:"counts"`
}

type UserWithBigModelsDTO struct {
	BigModels []UserWithBigModelDTO `bigmodels`
	Counts    int                   `json:"counts"`
}

type UsersDTO struct {
	Users  []domain.User `json:"users"`
	Counts int           `json:"counts"`
}

// d1Service struct
type d1Service struct {
	log *logrus.Entry
	ds  user.D1Service
}

func NewD1Service(
	log *logrus.Entry,
	ds user.D1Service,
) D1Service {
	d := &d1Service{
		log: log,
		ds:  ds,
	}

	return d
}

func (d d1Service) AddUserWithRepo(cmd *UserWithRepoCmd) error {
	return d.ds.AddUserWithRepo()
}

func (d d1Service) AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error {
	return nil
}

func (d d1Service) GetUserWithRepo() (UserWithRepoDTO, error) {
}

func (d d1Service) GetUserWithBigModel() (UserWithBigModelDTO, error) {
}

func (d d1Service) GetUsers() (UsersDTO, error) {
}

func (d d1Service) toUserWithRepoDTO(u *domain.UserWithRepo, dto *UserWithRepoDTO) {
	*dto = UserWithRepoDTO{
		Counts:   u.Counts,
		UpdateAt: u.UpdateAt,
	}
}

func (d d1Service) toUserWithBigModelDTO(u *domain.UserWithBigModel, dto *UserWithBigModelDTO) {
	*dto = UserWithBigModelDTO{
		BigModelType: u.BigModelType,
		Users:        u.Users,
		Counts:       u.Counts,
	}
}
