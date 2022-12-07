package main

import (
	"flag"
	"os"

	"github.com/opensourceways/community-robot-lib/logrusutil"
	liboptions "github.com/opensourceways/community-robot-lib/options"
	"github.com/sirupsen/logrus"

	// "github.com/opensourceways/xihe-training-center/app"
	// "github.com/opensourceways/xihe-training-center/controller"
	// "github.com/opensourceways/xihe-training-center/docs"
	// "github.com/opensourceways/xihe-training-center/domain"
	// "github.com/opensourceways/xihe-training-center/huaweicloud/trainingimpl"
	// "github.com/opensourceways/xihe-training-center/infrastructure/mysql"
	// "github.com/opensourceways/xihe-training-center/infrastructure/platformimpl"
	// "github.com/opensourceways/xihe-training-center/infrastructure/synclockimpl"
	// "github.com/opensourceways/xihe-training-center/infrastructure/watchimpl"
	// "github.com/opensourceways/xihe-training-center/server"
	"project/xihe-statistics/config"
	"project/xihe-statistics/infrastructure/mongodb"
	"project/xihe-statistics/server"
	"project/xihe-statistics/docs"
)

type options struct {
	service     liboptions.ServiceOptions
	enableDebug bool
}

func (o *options) Validate() error {
	return o.service.Validate()
}

func gatherOptions(fs *flag.FlagSet, args ...string) options {
	var o options

	o.service.AddFlags(fs)

	fs.BoolVar(
		&o.enableDebug, "enable_debug", false,
		"whether to enable debug model.",
	)

	fs.Parse(args)
	return o
}

func main() {
	logrusutil.ComponentInit("xihe-training-center")
	log := logrus.NewEntry(logrus.StandardLogger())

	o := gatherOptions(
		flag.NewFlagSet(os.Args[0], flag.ExitOnError),
		os.Args[1:]...,
	)
	if err := o.Validate(); err != nil {
		logrus.Fatalf("Invalid options, err:%s", err.Error())
	}

	if o.enableDebug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("debug enabled.")
	}

	// cfg
	var cfg string
	flag.StringVar(&cfg, "conf", "./conf/config.yaml", "指定配置文件路径")
	flag.Parse()
	// loading config file
	err := config.Init(cfg)
	if err != nil {
		panic(err)
	}

	// mongodb
	m := config.Conf.Mongodb
	if err := mongodb.Initialize(m.DBConn, m.DBName); err != nil {
		logrus.Fatalf("initialize mongodb failed, err:%s", err.Error())
	}

	defer mongodb.Close()

	// gin
	server.StartWebServer(docs.SwaggerInfo, &server.Service{
		Port:     o.service.Port,
		Timeout:  o.service.GracePeriod,
		Log:      log,
		Training: service,
	})
}
