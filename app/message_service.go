package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

// BigModel
type BigModelRecordMessageService interface {
	AddUserWithBigModel(*UserWithBigModelAddCmd) error
}

func NewBigModelRecordMessageService(
	ub repository.UserWithBigModel,
) BigModelRecordMessageService {
	return bigModelRecordMessageService{
		ub: ub,
	}
}

type bigModelRecordMessageService struct {
	ub repository.UserWithBigModel
}

func (b bigModelRecordMessageService) AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error { // implement app function with infrastructure function
	v := new(domain.UserWithBigModel)
	cmd.toBigModel(v)

	err := b.ub.Add(v)
	if err != nil {
		return err
	}
	return nil
}

// Repo
type RepoRecordMessageService interface {
	Add(*RepoRecordAddCmd) error
}

func NewRepoRecordMessageService(
	ur repository.UserWithRepo,
) RepoRecordMessageService {
	return repoRecordMessageService{
		ur: ur,
	}
}

type repoRecordMessageService struct {
	ur repository.UserWithRepo
}

func (s repoRecordMessageService) Add(cmd *RepoRecordAddCmd) (err error) {
	uwr := new(domain.UserWithRepo)
	cmd.toRepo(uwr)

	return s.ur.Add(uwr)
}

// Register
type RegisterRecordMessageService interface {
	Add(*domain.RegisterRecord) error
}

type registerRecordMessageService struct {
	rr repository.RegisterRecord
}

func NewRegisterRecordMessageService(
	rr repository.RegisterRecord,
) RegisterRecordMessageService {
	return registerRecordMessageService{
		rr: rr,
	}
}

func (s registerRecordMessageService) Add(d *domain.RegisterRecord) (err error) {
	return s.rr.Add(d)
}

// Download
type DownloadRecordMessageService interface {
	Add(*domain.DownloadRecord) error
}

func NewDownloadRecordMessgaeService(
	dr repository.DownloadRecord,
) DownloadRecordMessageService {
	return downloadRecordMessageService{
		dr: dr,
	}
}

type downloadRecordMessageService struct {
	dr repository.DownloadRecord
}

func (s downloadRecordMessageService) Add(d *domain.DownloadRecord) (err error) {
	return s.dr.Add(d)
}

// train
type TrainRecordMessageService interface {
	Add(*domain.TrainRecord) error
}

type trainRecordMessageService struct {
	tr repository.TrainRecord
}

func NewtrainRecordMessageService(
	tr repository.TrainRecord,
) TrainRecordMessageService {
	return trainRecordMessageService{
		tr: tr,
	}
}

func (s trainRecordMessageService) Add(tr *domain.TrainRecord) (err error) {
	return s.tr.Add(tr)
}
