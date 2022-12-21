package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/opensourceways/community-robot-lib/logrusutil"
	"github.com/sirupsen/logrus"

	"project/xihe-statistics/app"
	"project/xihe-statistics/config"
	"project/xihe-statistics/controller"
	"project/xihe-statistics/domain"
	"project/xihe-statistics/infrastructure/mongodb"
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

	// mongodb
	m := config.Conf.Mongodb
	if err := mongodb.Initialize(m.DBConn, m.DBName); err != nil {
		logrus.Fatalf("initialize mongodb failed, err:%s", err.Error())
	}

	defer mongodb.Close()

	// mq
	kafkaAddress := config.Conf.Message.KafKaAddress
	topic := config.Conf.Message.KafKaConfig.Topic
	go receive(kafkaAddress, topic) // TODO: goroutine must handle all msg even main process end

	// gin
	server.StartWebServer(config.Conf.HttpPort, time.Duration(config.Conf.Duration))
}

type msgBigModelCmd struct {
	UserName string `json:"username"`
	BigModel string `json:"bigmodel"`
}

func toCmd(msg string) (cmd app.UserWithBigModelAddCmd, err error) {
	var msgCmd msgBigModelCmd
	json.Unmarshal([]byte(msg), &msgCmd)
	fmt.Printf("msgCmd: %v\n", msgCmd)
	if cmd.BigModel, err = domain.NewBigModel(msgCmd.BigModel); err != nil {
		return
	}

	cmd.UserName = msgCmd.UserName

	return
}

func bigModelTask(msg string) error {
	// mq
	collections := config.Conf.Mongodb.MongodbCollections

	bigModelRecord := repositories.NewBigModelRecordRepository(
		mongodb.NewBigModelMapper(collections.BigModel),
	)

	service := app.NewBigModelRecordService(bigModelRecord)

	cmd, err := toCmd(msg)
	if err != nil {
		return err
	}
	service.AddUserWithBigModel(&cmd)

	return nil
}

func receive(address string, topic string) {
	//配置
	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置kafka版本号
	config.Version = sarama.V3_1_0_0
	//新建一个消费者
	consumer, err := sarama.NewConsumer([]string{address}, config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic("create comsumer failed")
	}
	defer consumer.Close()
	//特定分区消费者，需要设置主题，分区和偏移量，sarama.OffsetNewest表示每次从最新的消息开始消费
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Println("error get partition sonsumer")
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			err2 := bigModelTask(string(msg.Value))
			if err2 != nil {
				panic(err2)
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Err)
		}
	}
}

// func run(d *messages.BigModel, log *logrus.Entry) {
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

// 	var wg sync.WaitGroup
// 	defer wg.Wait()

// 	called := false
// 	ctx, done := context.WithCancel(context.Background())

// 	defer func() {
// 		if !called {
// 			called = true
// 			done()
// 		}
// 	}()

// 	wg.Add(1)
// 	go func(ctx context.Context) {
// 		defer wg.Done()

// 		select {
// 		case <-ctx.Done():
// 			log.Info("receive done. exit normally")
// 			return

// 		case <-sig:
// 			log.Info("receive exit signal")
// 			done()
// 			called = true
// 			return
// 		}
// 	}(ctx)

// 	if err := d.Run(ctx, log); err != nil {
// 		log.Errorf("subscribe failed, err:%v", err)
// 	}
// }
