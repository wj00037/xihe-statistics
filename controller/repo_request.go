package controller

import (
	"project/xihe-statistics/app"
)

type AddRepoRecordRequest struct {
	UserName string `json:"username"`
	RepoName string `json:"repo_name"`
}

func (req *AddRepoRecordRequest) toCmd() (cmd app.RepoRecordAddCmd, err error) {
	cmd.UserName = req.UserName
	cmd.RepoName = req.RepoName
	return
}
