package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/opensourceways/community-robot-lib/logrusutil"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/app"
	"project/xihe-statistics/config"
	"project/xihe-statistics/controller"
	"project/xihe-statistics/infrastructure/messages"
	"project/xihe-statistics/infrastructure/pgsql"
	"project/xihe-statistics/infrastructure/repositories"
	"project/xihe-statistics/server"
)

func main() {
	logrusutil.ComponentInit("xihe-statistics")
	log := logrus.NewEntry(logrus.StandardLogger())

	// cfg
	var cfg string
	flag.StringVar(&cfg, "conf", "./conf/config.yaml", "指定配置文件路径")
	flag.Parse()
	// loading config file
	err := config.Init(cfg)
	if err != nil {
		panic(err)
	}

	// controller
	controller.Init(log)

	// pgsql
	if err := pgsql.Initialize(config.Conf.PGSQL); err != nil {
		logrus.Fatalf("initialize pgsql failed, err:%s", err.Error())
	}

	// init kafka
	if err := messages.Init(config.Conf.GetMQConfig(), log, config.Conf.Topics); err != nil {
		log.Fatalf("initialize mq failed, err:%v", err)
	}

	defer messages.Exit(log)

	go run(newHandler(config.Conf, log), log)

	// gin
	server.StartWebServer(config.Conf.HttpPort, time.Duration(config.Conf.Duration))
}

func newHandler(cfg *config.SrvConfig, log *logrus.Entry) *messages.Handler {

	bigModelRecord := repositories.NewBigModelRecordRepository(
		// infrastructure.mongodb -> infrastructure.repositories (mapper)
		pgsql.NewBigModelMapper(pgsql.BigModelRecord{}),
	)

	repoRecord := repositories.NewUserWithRepoRepository(
		pgsql.NewUserWithRepoMapper(pgsql.UserWithRepo{}),
	)

	bs := app.NewBigModelRecordMessageService(bigModelRecord)
	rs := app.NewRepoRecordMessageService(repoRecord)

	return &messages.Handler{
		Log:      log,
		MaxRetry: config.Conf.MaxRetry,

		BigModel: bs,
		Repo:     rs,
	}
}

func run(h *messages.Handler, log *logrus.Entry) {
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

	if err := messages.Subscribe(ctx, h, log); err != nil {
		log.Errorf("subscribe failed, err:%v", err)
	}
}
