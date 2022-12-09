package app

// import (
// 	"errors"

// 	"github.com/sirupsen/logrus"

// 	"project/xihe-statistics/domain"
// 	"project/xihe-statistics/domain/user"
// )

// func (cmd *UserWithRepoCmd) Validate() error {
// 	err := errors.New("invalid cmd of add user create repo record")

// 	b := cmd.UserName != ""

// 	if !b {
// 		return err
// 	}

// 	return nil

// }

// func (cmd *UserWithBigModelAddCmd) Validate() error {
// 	err := errors.New("invalid cmd of add user query big model record")

// 	b := cmd.UserName != "" &&
// 		cmd.BigModel != ""

// 	if !b {
// 		return err
// 	}

// 	t := cmd.BigModel == "taichu-VQA" ||
// 		cmd.BigModel == "taichu-TextToImg" ||
// 		cmd.BigModel == "taichu-ImgToText"
// 	if !t {
// 		return err
// 	}

// 	return nil
// }

// type D1Service interface {
// 	AddUserWithRepo(cmd *UserWithRepoCmd) error
// 	AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error
// 	GetUserWithRepoCounts() (UserWithRepoCountsDTO, error)
// 	GetUserWithBigModelCounts() (UserWithBigModelCountsDTO, error)
// 	GetUsers() (UsersDTO, error)
// }

// type UserWithRepoCmd struct {
// 	domain.UserWithRepo
// }

// type UserWithBigModelAddCmd struct {
// 	domain.UserWithBigModel
// }

// type UserWithRepoCountsDTO struct {
// 	Counts   int   `json:"counts"`
// 	UpdateAt int64 `json:"update_at"`
// }

// type UserWithBigModelCountsDTO struct {
// 	BigModelType string        `json:"bigmodel_type"`
// 	Users        []domain.User `json:"users"`
// 	Counts       int           `json:"counts"`
// }

// type UsersDTO struct {
// 	Users  []domain.User `json:"users"`
// 	Counts int           `json:"counts"`
// }

// // d1Service struct
// type d1Service struct {
// 	log *logrus.Entry
// 	ur  user.UserWithRepo
// 	um  user.UserWithBigModel
// }

// func NewD1Service(
// 	log *logrus.Entry,
// 	ur user.UserWithRepo,
// 	um user.UserWithBigModel,
// ) D1Service {
// 	d := &d1Service{
// 		log: log,
// 		ur:  ur,
// 		um:  um,
// 	}

// 	return d
// }

// func (d d1Service) AddUserWithRepo(cmd *UserWithRepoCmd) error {
// 	return d.ur.Add(&cmd.UserWithRepo)
// }

// func (d d1Service) AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error {
// 	return d.um.Add(&cmd.UserWithBigModel)
// }

// func (d d1Service) GetUserWithRepoCounts() (UserWithRepoDTO, error) {
// }

// func (d d1Service) GetUserWithBigModelCounts() (UserWithBigModelDTO, error) {
// }

// func (d d1Service) GetUsers() (UsersDTO, error) {
// }

// func (d d1Service) toUserWithRepoCountsDTO(u *domain.UserWithRepo, dto *UserWithRepoDTO) {
// 	*dto = UserWithRepoDTO{
// 		Counts:   u.Counts,
// 		UpdateAt: u.UpdateAt,
// 	}
// }

// func (d d1Service) toUserWithBigModelCountsDTO(u *domain.UserWithBigModel, dto *UserWithBigModelDTO) {
// 	*dto = UserWithBigModelDTO{
// 		BigModelType: u.BigModelType,
// 		Users:        u.Users,
// 		Counts:       u.Counts,
// 	}
// }
