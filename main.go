package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/opensourceways/community-robot-lib/kafka"
	"github.com/opensourceways/community-robot-lib/logrusutil"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/config"
	"project/xihe-statistics/controller"
	"project/xihe-statistics/infrastructure/messages"
	"project/xihe-statistics/infrastructure/mongodb"
	"project/xihe-statistics/server"
)

// type options struct {
// 	service     liboptions.ServiceOptions
// 	enableDebug bool
// }

// func (o *options) Validate() error {
// 	return o.service.Validate()
// }

// func gatherOptions(fs *flag.FlagSet, args ...string) options {
// 	var o options

// 	o.service.AddFlags(fs)

// 	fs.BoolVar(
// 		&o.enableDebug, "enable_debug", false,
// 		"whether to enable debug model.",
// 	)

// 	fs.Parse(args)
// 	return o
// }

func main() {
	logrusutil.ComponentInit("xihe-statistics")
	log := logrus.NewEntry(logrus.StandardLogger())

	// o := gatherOptions(
	// 	flag.NewFlagSet(os.Args[0], flag.ExitOnError),
	// 	os.Args[1:]...,
	// )
	// if err := o.Validate(); err != nil {
	// 	logrus.Fatalf("Invalid options, err:%s", err.Error())
	// }

	// if o.enableDebug {
	// 	logrus.SetLevel(logrus.DebugLevel)
	// 	logrus.Debug("debug enabled.")
	// }

	// cfg
	var cfg string
	flag.StringVar(&cfg, "conf", "./conf/config.yaml", "指定配置文件路径")
	flag.Parse()
	// loading config file
	err := config.Init(cfg)
	if err != nil {
		panic(err)
	}

	// kafka
	kafkaCfg, err := messages.LoadKafkaConfig(config.Conf.Message.KafKaConfigFile)
	if err != nil {
		log.Errorf("Error loading kfk config, err:%v", err)

		return
	}

	if err = messages.ConnectKafKa(&kafkaCfg); err != nil {
		log.Errorf("Error connecting kfk mq, err:%v", err)

		return
	}

	defer kafka.Disconnect()

	// controller
	controller.Init(log)

	// mongodb
	m := config.Conf.Mongodb
	if err := mongodb.Initialize(m.DBConn, m.DBName); err != nil {
		logrus.Fatalf("initialize mongodb failed, err:%s", err.Error())
	}

	defer mongodb.Close()

	// mq

	// gin
	server.StartWebServer(config.Conf.HttpPort, time.Duration(config.Conf.Duration))
}

func run(d *messages.BigModel, log *logrus.Entry) {
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

	if err := d.Run(ctx, log); err != nil {
		log.Errorf("subscribe failed, err:%v", err)
	}
}
