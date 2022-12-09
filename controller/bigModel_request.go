package controller

import (
	"fmt"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
	"time"
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

	// add timestamp
	cmd.CreateAt = time.Now().Unix() // TODO: delete time.now

	fmt.Printf("cmd: %v\n", cmd)

	return

}
