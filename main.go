package main

import (
	"context"
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/opensourceways/community-robot-lib/kafka"
	"github.com/opensourceways/community-robot-lib/logrusutil"
	"github.com/opensourceways/community-robot-lib/mq"
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

	// // mongodb
	// m := config.Conf.Mongodb
	// if err := mongodb.Initialize(m.DBConn, m.DBName); err != nil {
	// 	logrus.Fatalf("initialize mongodb failed, err:%s", err.Error())
	// }

	// defer mongodb.Close()

	// init kafka
	kafkaCfg, err := loadKafkaConfig(config.Conf.KafKaConfigFile)
	if err != nil {
		log.Errorf("Error loading kfk config, err:%v", err)

		return
	}

	if err := connetKafka(&kafkaCfg); err != nil {
		log.Errorf("Error connecting kfk mq, err:%v", err)

		return
	}

	defer kafka.Disconnect()

	bigModelRecord := repositories.NewBigModelRecordRepository(
		// infrastructure.mongodb -> infrastructure.repositories (mapper)
		pgsql.NewBigModelMapper(pgsql.BigModelRecord{}),
	)

	// repoRecord := repositories.NewUserWithRepoRepository(
	// 	mongodb.NewUserWithRepoMapper(collections.Repo),
	// )

	repoRecord := repositories.NewUserWithRepoRepository(
		pgsql.NewUserWithRepoMapper(pgsql.UserWithRepo{}),
	)

	bs := app.NewBigModelRecordService(bigModelRecord)
	rs := app.NewRepoRecordService(repoRecord)

	d := messages.NewAddRecord(config.Conf.MQConfig, bs, rs)

	go run(d, log)

	// gin
	server.StartWebServer(config.Conf.HttpPort, time.Duration(config.Conf.Duration))
}

func connetKafka(cfg *mq.MQConfig) error {
	tlsConfig, err := cfg.TLSConfig.TLSConfig()
	if err != nil {
		return err
	}

	err = kafka.Init(
		mq.Addresses(cfg.Addresses...),
		mq.SetTLSConfig(tlsConfig),
		mq.Log(logrus.WithField("module", "kfk")),
	)
	if err != nil {
		return err
	}

	return kafka.Connect()
}

func loadKafkaConfig(file string) (cfg mq.MQConfig, err error) {
	v, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	str := string(v)
	if str == "" {
		err = errors.New("missing addresses")

		return
	}

	addresses := parseAddress(str)
	if len(addresses) == 0 {
		err = errors.New("no valid address for kafka")

		return
	}

	if err = kafka.ValidateConnectingAddress(addresses); err != nil {
		return
	}

	cfg.Addresses = addresses

	return
}

func parseAddress(addresses string) []string {
	var reIpPort = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}:[1-9][0-9]*$`)

	v := strings.Split(addresses, ",")
	r := make([]string, 0, len(v))
	for i := range v {
		if reIpPort.MatchString(v[i]) {
			r = append(r, v[i])
		}
	}

	return r
}

func run(d *messages.AddRecord, log *logrus.Entry) {
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
