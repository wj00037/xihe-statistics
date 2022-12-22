package messages

import (
	"encoding/json"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/infrastructure/pgsql"
	"project/xihe-statistics/infrastructure/repositories"
)

type msgBigModelCmd struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	BigModel string `json:"bigmodel"`
	CreateAt int64  `json:"create_at"`
}

func toBigModelCmd(msg string) (cmd app.UserWithBigModelAddCmd, err error) {
	var msgCmd msgBigModelCmd
	json.Unmarshal([]byte(msg), &msgCmd)
	if cmd.BigModel, err = domain.NewBigModel(msgCmd.BigModel); err != nil {
		return
	}

	cmd.UserName = msgCmd.UserName
	cmd.CreateAt = msgCmd.CreateAt

	return
}

func bigModelTask(msg string) error {
	// mq
	bigModelRecord := repositories.NewBigModelRecordRepository(
		// infrastructure.mongodb -> infrastructure.repositories (mapper)
		pgsql.NewBigModelMapper(pgsql.BigModelRecord{}),
	)

	service := app.NewBigModelRecordService(bigModelRecord)

	cmd, err := toBigModelCmd(msg)
	if err != nil {
		return err
	}

	service.AddUserWithBigModel(&cmd)

	return nil
}

type msgRepoCmd struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	RepoName string `json:"repo_name"`
	CreateAt int64  `json:"create_at"`
}

func (m msgRepoCmd) toRepoCmd(msg string) (cmd app.RepoRecordAddCmd, err error) {
	json.Unmarshal([]byte(msg), &m)

	cmd.RepoName = m.RepoName
	cmd.UserName = m.UserName
	cmd.CreateAt = m.CreateAt

	return
}

func repoTask(msg string) error {
	// mq
	repoRecord := repositories.NewUserWithRepoRepository(
		pgsql.NewUserWithRepoMapper(pgsql.UserWithRepo{}),
	)

	service := app.NewRepoRecordService(repoRecord)

	var m msgRepoCmd
	cmd, err := m.toRepoCmd(msg)
	if err != nil {
		return err
	}

	service.Add(&cmd)

	return nil
}
