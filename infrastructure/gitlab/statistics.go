package gitlab

import (
	"project/xihe-statistics/app"
	"project/xihe-statistics/config"
	"project/xihe-statistics/infrastructure/pgsql"
	"project/xihe-statistics/infrastructure/repositories"
	"time"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Log *logrus.Entry

	gs app.GitLabService
}

func NewHandler(cfg *config.Config, log *logrus.Entry) *Handler {
	platform := NewGitlabStatistics(cfg)

	gitlabRecord := repositories.NewGitLabRecordRepository(
		pgsql.NewGitLabRecordMapper(pgsql.GitLabRecord{}),
	)
	gs := app.NewGitLabService(platform, gitlabRecord)

	return &Handler{
		Log: log,

		gs: gs,
	}
}

func Do(h *Handler) error {
	dto, err := h.gs.Counts()
	if err != nil {
		return err
	}

	// TODO: must check if dto."counts" is bigger than before,
	// in case some project deleted makes counts decreasing
	cmd := app.CloneCountsCmd(dto)
	return h.gs.Save(&cmd)
}

func Run(h *Handler, log *logrus.Entry, cfg *config.Config) {
	ticker := time.NewTicker(time.Second * 86400)	// TODO: check config reading
	defer ticker.Stop()

	// to excute Do at function Run started
	err := Do(h)
	if err != nil {
		log.Errorf("count clone failed, err:%v", err)
	}

	for range ticker.C {
		err := Do(h)
		if err != nil {
			log.Errorf("count clone failed, err:%v", err)
		}
	}
}
