package controller

import (
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
)

type QueryBigModelRequest struct {
	UserName string `json:"username"`
	BigModel string `json:"bigmodel"`
}

func (req *QueryBigModelRequest) toCmd() (cmd app.UserWithBigModelAddCmd, err error) {

	if cmd.BigModel, err = domain.NewBigModel(req.BigModel); err != nil {
		return
	}

	cmd.UserName = req.UserName

	return
}
