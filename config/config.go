package config

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(SrvConfig)

type SrvConfig struct {
	Name            string `mapstructure:"name"`
	HttpPort        int    `mapstructure:"http_port"`
	Duration        int    `mapstructure:"duration"`
	KafKaConfigFile string `mapstructure:"kafka_config_file"`
	*MQConfig       `mapstructure:"mq_config"`
	*PGSQL          `mapstructure:"pgsql"`
	*Mongodb        `mapstructure:"mongodb"`
	*Message        `mapstructure:"message"`
}

type MQConfig struct {
	// AccessEndpoint is used to send back the message.
	AccessEndpoint string `mapstructure:"access_endpoint"  required:"true"`
	AccessHmac     string `mapstructure:"access_hmac"      required:"true"`

	Topic     string `mapstructure:"topic"                 required:"true"`
	UserAgent string `mapstructure:"user_agent"            required:"true"`

	// The unit is Gbyte
	SizeOfWorspace int `mapstructure:"size_of_workspace"   required:"true"`

	// The unit is Gbyte
	AverageRepoSize int `mapstructure:"average_repo_size"  required:"true"`
}

func (cfg *MQConfig) ConcurrentSize() int {
	return cfg.SizeOfWorspace / (cfg.AverageRepoSize) / 2
}

func (cfg *MQConfig) Validate() error {
	if cfg.Topic == "" {
		return errors.New("missing topic")
	}

	if cfg.UserAgent == "" {
		return errors.New("missing user_agent")
	}

	if cfg.AverageRepoSize <= 0 {
		return errors.New("invalid average_repo_size")
	}

	if cfg.ConcurrentSize() <= 0 {
		return errors.New("the concurrent size <= 0")
	}

	return nil
}

type PGSQL struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"db_name"`
	Password string `mapstructure:"password"`
}

type Mongodb struct {
	DBName              string `mapstructure:"db_name"`
	DBConn              string `mapstructure:"db_conn"`
	*MongodbCollections `mapstructure:"collections"`
}

type MongodbCollections struct {
	D0       string `mapstructure:"d0"`
	BigModel string `mapstructure:"bigmodel"`
	Repo     string `mapstructure:"repo"`
	D1       string `mapstructure:"d1"`
	D2       string `mapstructure:"d2"`
}

type Message struct {
	KafKaAddress string `mapstructure:"kafka_address"`
	*KafKaConfig `mapstructure:"kafka_config"`
}

type KafKaConfig struct {
	// AccessEndpoint is used to send back the message.
	AccessEndpoint string `mapstructure:"access_endpoint"`
	AccessHmac     string `mapstructure:"access_hmac"`

	Topic     string `mapstructure:"topic"`
	UserAgent string `mapstructure:"user_agent"`

	// The unit is Gbyte
	SizeOfWorspace int `mapstructure:"size_of_workspace"`

	// The unit is Gbyte
	AverageRepoSize int `mapstructure:"average_repo_size"`
}

// Init 整个服务配置文件初始化的方法
func Init(filePath string) (err error) {
	// 方式1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	// viper.SetConfigFile("./conf/config.yaml")
	viper.SetConfigFile(filePath)

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}
	// 如果使用的是 viper.GetXxx()方式使用配置的话，就无须下面的操作

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig() // 配置文件监听
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}

func (cfg *KafKaConfig) ConcurrentSize() int {
	return cfg.SizeOfWorspace / (cfg.AverageRepoSize) / 2
}

func (cfg *KafKaConfig) Validate() error {
	if cfg.Topic == "" {
		return errors.New("missing topic")
	}

	if cfg.UserAgent == "" {
		return errors.New("missing user_agent")
	}

	if cfg.AverageRepoSize <= 0 {
		return errors.New("invalid average_repo_size")
	}

	if cfg.ConcurrentSize() <= 0 {
		return errors.New("the concurrent size <= 0")
	}

	return nil
}
