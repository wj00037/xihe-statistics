package main

import (
	"flag"
	"os"
	"time"

	"github.com/opensourceways/community-robot-lib/logrusutil"
	liboptions "github.com/opensourceways/community-robot-lib/options"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/config"
	"project/xihe-statistics/controller"
	"project/xihe-statistics/infrastructure/gitlab"
	"project/xihe-statistics/infrastructure/messages"
	"project/xihe-statistics/infrastructure/pgsql"
	"project/xihe-statistics/server"
)

type options struct {
	service     liboptions.ServiceOptions
	enableDebug bool
}

func (o *options) Validate() error {
	return o.service.Validate()
}

func gatherOptions(fs *flag.FlagSet, args ...string) (options, error) {
	var o options

	o.service.AddFlags(fs)

	fs.BoolVar(
		&o.enableDebug, "enable_debug", false,
		"whether to enable debug model.",
	)

	err := fs.Parse(args)

	return o, err
}

func main() {
	logrusutil.ComponentInit("xihe-statistics")
	log := logrus.NewEntry(logrus.StandardLogger())

	o, err := gatherOptions(
		flag.NewFlagSet(os.Args[0], flag.ExitOnError),
		os.Args[1:]...,
	)
	if err != nil {
		logrus.Fatalf("new options failed, err:%s", err.Error())
	}

	if err := o.Validate(); err != nil {
		logrus.Fatalf("Invalid options, err:%s", err.Error())
	}

	if o.enableDebug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("debug enabled.")
	}

	// cfg
	cfg := new(config.Config)
	if err := config.LoadConfig(o.service.ConfigFile, cfg); err != nil {
		logrus.Fatalf("load config, err:%s", err.Error())
	}

	if err := os.Remove(o.service.ConfigFile); err != nil {
		logrus.Fatalf("remove config file failed, err:%s", err.Error())
	}

	// controller
	controller.Init(log)

	// pgsql
	if err := pgsql.Initialize(&cfg.PGSQL); err != nil {
		logrus.Fatalf("initialize pgsql failed, err:%s", err.Error())
	}

	if err := os.Remove(cfg.PGSQL.DBCert); err != nil {
		logrus.Fatalf("postgresql dbcert file delete failed, err:%s", err.Error())
	}

	// init kafka
	if err := messages.Init(cfg.GetKfkConfig(), log, cfg.MQTopics); err != nil {
		log.Fatalf("initialize mq failed, err:%v", err)
	}

	defer messages.Exit(log)

	// mq
	go messages.Run(messages.NewHandler(cfg, log), log)

	// gitlab statisitc
	go gitlab.Run(gitlab.NewHandler(cfg, log), log, cfg)

	// gin
	server.StartWebServer(cfg.HttpPort, time.Duration(cfg.Duration), cfg)
}
