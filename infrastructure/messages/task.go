package messages

import (
	"encoding/json"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain"
	"strings"
)

// type addRecordTask struct {
// 	msgBigModelCmd
// 	msgRepoCmd
// }

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

type msgRepoCmd struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	RepoName string `json:"repo_name"`
	CreateAt int64  `json:"create_at"`
}

func toRepoCmd(msg string) (cmd app.RepoRecordAddCmd, err error) {
	var msgCmd msgRepoCmd

	json.Unmarshal([]byte(msg), &msgCmd)

	cmd.RepoName = msgCmd.RepoName
	cmd.UserName = msgCmd.UserName
	cmd.CreateAt = msgCmd.CreateAt

	return
}

type addRecordTaskGenerator struct {
}

func (d *addRecordTaskGenerator) genTask(payload []byte, header map[string]string) (
	cmd interface{}, ok bool, err error,
) {
	if find := strings.Contains(string(payload), "statistics-bigmodel"); find { // TODO: simplify
		cmd, err = toBigModelCmd(string(payload))
		if err != nil {
			ok = false
			return
		}
	}

	if find := strings.Contains(string(payload), "statistics-repo"); find {
		cmd, err = toRepoCmd(string(payload))
		if err != nil {
			ok = false
			return
		}
	}

	ok = true

	return

}
