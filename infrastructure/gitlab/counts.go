package gitlab

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"project/xihe-statistics/config"
	"project/xihe-statistics/domain/platform"

	"github.com/opensourceways/community-robot-lib/utils"
)

func NewGitlabStatistics(cfg *config.SrvConfig) platform.PlatForm {
	return &gitlabStatistics{
		token:        cfg.GitLab.RootToken,
		endpoint:     cfg.GitLab.Endponit,
		countPerPage: cfg.CountPerPage,
		cli:          utils.NewHttpClient(3),
	}
}

type gitlabStatistics struct {
	token        string
	endpoint     string
	countPerPage int
	cli          utils.HttpClient
}

func (impl *gitlabStatistics) GetProjectId(pageNum int) ([]platform.ProjectId, error) {
	return impl.getProjectId(pageNum)
}

func (impl *gitlabStatistics) GetCloneTotal(id int) (total platform.CloneTotal, err error) {
	res, err := impl.getCloneTotal(id)
	if err != nil {
		return
	}
	return res.toCloneTotal(), nil
}

func (impl *gitlabStatistics) getProjectId(pageNum int) (resp []platform.ProjectId, err error) {
	url := fmt.Sprintf("%s/projects/?simple=true&per_page=%d&page=%d", impl.endpoint, impl.countPerPage, pageNum)
	req, err := impl.newRequest(impl.token, url, "GET", nil)
	if err != nil {
		return
	}

	_, err = impl.cli.ForwardTo(req, &resp)
	if err != nil {
		return
	}

	return
}

func (impl *gitlabStatistics) getCloneTotal(id int) (resp cloneTotalResult, err error) {
	url := fmt.Sprintf("%s/projects/%d/statistics", impl.endpoint, id)
	req, err := impl.newRequest(impl.token, url, "GET", nil)
	if err != nil {
		return
	}

	_, err = impl.cli.ForwardTo(req, &resp)
	if err != nil {
		return
	}

	return
}

func (impl *gitlabStatistics) newRequest(
	token, url, method string, param interface{},
) (*http.Request, error) {
	var body io.Reader

	if param != nil {
		v, err := utils.JsonMarshal(param)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(v)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	h := &req.Header
	h.Add("PRIVATE-TOKEN", token)
	h.Add("Content-Type", "application/json")

	return req, nil
}
