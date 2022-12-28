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
	Name     string `mapstructure:"name"`
	HttpPort int    `mapstructure:"http_port"`
	Duration int    `mapstructure:"duration"`
	*PGSQL   `mapstructure:"pgsql"`
	*MQ      `mapstructure:"mq"`
}

type PGSQL struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"db_name"`
	Password string `mapstructure:"password"`
}

type MQ struct {
	Address  string `mapstructure:"address"`
	MaxRetry int    `mapstructure:"max_retry"`
	Topics   `mapstructure:"topics"`
}

type Topics struct {
	Statistics string `mapstructure:"statistics" json:"statistics" required:"true"`
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

// viper init
func Init(filePath string) (err error) {
	viper.SetConfigFile(filePath)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
