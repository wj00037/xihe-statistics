package messages

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
)

func Receive(address string, topic string) {
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
			if find := strings.Contains(string(msg.Value), "statistics-bigmodel"); find { // TODO: simplify
				err = bigModelTask(string(msg.Value))
				if err != nil {
					panic(err)
				}
			}

			if find := strings.Contains(string(msg.Value), "statistics-repo"); find {
				err = repoTask(string(msg.Value))
				if err != nil {
					panic(err)
				}
			}

		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Err)
		}
	}
}
