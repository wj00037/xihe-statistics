package messages

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"

	"github.com/opensourceways/community-robot-lib/kafka"
	"github.com/opensourceways/community-robot-lib/mq"
	"github.com/sirupsen/logrus"
)

func ConnectKafKa(cfg *mq.MQConfig) error {
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

func Exit(log *logrus.Entry) {
	if err := kafka.Disconnect(); err != nil {
		log.Error("exit kafka, err:%v", err)
	}
}

func LoadKafkaConfig(file string) (cfg mq.MQConfig, err error) {
	v, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	str := string(v)
	if str == "" {
		err = errors.New("missing addresses")

		return
	}

	addresses := ParseAddress(str)
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

func ParseAddress(addresses string) []string {
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

func Run(d *syncrepo.SyncRepo, log *logrus.Entry) {
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
