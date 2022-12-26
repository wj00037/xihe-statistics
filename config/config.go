package config

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/opensourceways/community-robot-lib/mq"
	"github.com/spf13/viper"
)

var reIpPort = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}:[1-9][0-9]*$`)
var Conf = new(SrvConfig)

type SrvConfig struct {
	Name            string `mapstructure:"name"`
	HttpPort        int    `mapstructure:"http_port"`
	Duration        int    `mapstructure:"duration"`
	KafKaConfigFile string `mapstructure:"kafka_config_file"`
	MaxRetry        int    `mapstructure:"max_retry"`
	*PGSQL          `mapstructure:"pgsql"`
	*Mongodb        `mapstructure:"mongodb"`
	*MQ             `mapstructure:"mq"`
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

type MQ struct {
	Address string `mapstructure:"address"`
	Topics  `mapstructure:"topics"`
}

type Topics struct {
	StatisticsBigModel string `mapstructure:"statistics_bigmodel" json:"statistics_bigmodel" required:"true"`
	StatisticsRepo     string `mapstructure:"statistics_repo" json:"statistics_repo" required:"true"`
}

func (cfg *SrvConfig) GetMQConfig() mq.MQConfig {
	return mq.MQConfig{
		Addresses: cfg.MQ.ParseAddress(),
	}
}

func (cfg *MQ) Validate() error {
	if r := cfg.ParseAddress(); len(r) == 0 {
		return errors.New("invalid mq address")
	}

	return nil
}

func (cfg *MQ) ParseAddress() []string {
	v := strings.Split(cfg.Address, ",")
	r := make([]string, 0, len(v))
	for i := range v {
		if reIpPort.MatchString(v[i]) {
			r = append(r, v[i])
		}
	}

	return r
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
