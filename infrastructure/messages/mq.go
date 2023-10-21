package messages

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	kfklib "github.com/opensourceways/kafka-lib/agent"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/app"
	"project/xihe-statistics/config"
	"project/xihe-statistics/infrastructure/pgsql"
	"project/xihe-statistics/infrastructure/repositories"
)

func Init(cfg kfklib.Config, log *logrus.Entry, topic config.Topics) error {
	topics = topic

	return kfklib.Init(&cfg, log, nil, "", true)
}

func Exit(log *logrus.Entry) {
	kfklib.Exit()
}

func NewHandler(cfg *config.Config, log *logrus.Entry) *Handler {

	bigModelRecord := repositories.NewBigModelRecordRepository(
		// infrastructure.mongodb -> infrastructure.repositories (mapper)
		pgsql.NewBigModelMapper(pgsql.BigModelRecord{}),
	)

	repoRecord := repositories.NewUserWithRepoRepository(
		pgsql.NewUserWithRepoMapper(pgsql.UserWithRepo{}),
	)

	registerRecord := repositories.NewRegisterRecordRepository(
		pgsql.NewRegisterRecordMapper(pgsql.RegisterRecord{}),
	)

	fileUploadRecord := repositories.NewFileUploadRecordRepository(
		pgsql.NewFileUploadRecordMapper(pgsql.FileUploadRecord{}),
	)

	downloadRecord := repositories.NewDownloadRecordRepository(
		pgsql.NewDownloadRecordMapper(pgsql.DownloadRecord{}),
	)

	trainRecord := repositories.NewTrainRecordRepository(
		pgsql.NewTrainRecordMapper(pgsql.TrainRecord{}),
	)

	cloudRecord := repositories.NewCloudRecordRepository(
		pgsql.NewCloudRecordMapper(pgsql.CloudRecord{}),
	)

	gitlabRecord := repositories.NewGitLabRecordRepository(
		pgsql.NewGitLabRecordMapper(pgsql.GitLabRecord{}),
	)

	bs := app.NewBigModelRecordService(bigModelRecord)
	rs := app.NewRepoRecordService(repoRecord)
	rr := app.NewRegisterRecordService(registerRecord)
	fr := app.NewFileUploadRecordService(fileUploadRecord)
	ds := app.NewDownloadRecordService(downloadRecord, gitlabRecord)
	ts := app.NewTrainRecordService(trainRecord)
	cr := app.NewCloudRecodeService(cloudRecord)

	return &Handler{
		Log: log,

		BigModel:   bs,
		Repo:       rs,
		Register:   rr,
		FileUpload: fr,
		Download:   ds,
		Train:      ts,
		Cloud:      cr,
	}
}

func Run(h *Handler, log *logrus.Entry) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	defer wg.Wait()

	called := false
	ctx, done := context.WithCancel(context.Background())

	defer func() {
		if !called {
			called = true
			done()
		}
	}()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		select {
		case <-ctx.Done():
			log.Info("receive done. exit normally")
			return

		case <-sig:
			log.Info("receive exit signal")
			done()
			called = true
			return
		}
	}(ctx)

	if err := Subscribe(ctx, h, log); err != nil {
		log.Errorf("subscribe failed, err:%v", err)
	}
}
