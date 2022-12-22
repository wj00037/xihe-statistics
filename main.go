package main

import (
	"flag"
	"time"

	"github.com/opensourceways/community-robot-lib/logrusutil"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/config"
	"project/xihe-statistics/controller"
	"project/xihe-statistics/infrastructure/messages"
	"project/xihe-statistics/infrastructure/pgsql"
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

	// // mongodb
	// m := config.Conf.Mongodb
	// if err := mongodb.Initialize(m.DBConn, m.DBName); err != nil {
	// 	logrus.Fatalf("initialize mongodb failed, err:%s", err.Error())
	// }

	// defer mongodb.Close()

	// mq
	kafkaAddress := config.Conf.Message.KafKaAddress
	topic := config.Conf.Message.KafKaConfig.Topic
	go messages.Receive(kafkaAddress, topic) // TODO: goroutine must handle all msg even main process end

	// gin
	server.StartWebServer(config.Conf.HttpPort, time.Duration(config.Conf.Duration))
}
