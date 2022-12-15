package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(SrvConfig)

type SrvConfig struct {
	Name     string `mapstructure:"name"`
	HttpPort int    `mapstructure:"http_port"`
	Duration int    `mapstructure:"duration"`
	*Mongodb `mapstructure:"mongodb"`
	*Message `mapstructure:"message"`
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
	KafKaConfig string `mapstructure:"kafka_config"`
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
