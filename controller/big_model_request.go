package controller

import (
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
)

type BigModelCreateRequest struct {
	UserName string `json:"username"`
	BigModel string `json:"bigmodel"`
}

func (req *BigModelCreateRequest) toCmd() (cmd app.UserWithBigModelAddCmd, err error) {

	if cmd.BigModel, err = domain.NewBigModel(req.BigModel); err != nil {
		return
	}

	if cmd.UserName, err = domain.NewAccount(req.UserName); err != nil {
		return
	}

	return
}

type BigModelQueryWithTypeAndTimeRequest struct {
	BigModel  string `json:"bigmodel"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (req BigModelQueryWithTypeAndTimeRequest) toCmd() (
	cmd app.BigModelCountIncreaseCmd,
	err error,
) {
	if cmd.BigModel, err = domain.NewBigModel(req.BigModel); err != nil {
		return
	}

	cmd.StartTime = req.StartTime
	cmd.EndTime = req.EndTime

	return
}
