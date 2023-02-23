package config

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/opensourceways/community-robot-lib/mq"
	"github.com/opensourceways/community-robot-lib/utils"
)

var reIpPort = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}:[1-9][0-9]*$`)

func LoadConfig(path string, cfg interface{}) error {
	if err := utils.LoadFromYaml(path, cfg); err != nil {
		return err
	}

	return nil
}

type Config struct {
	Name     string `json:"name"`
	HttpPort int    `json:"http_port"`
	Duration int    `json:"duration"`
	PGSQL    PGSQL  `json:"pgsql"`
	MQ       MQ     `json:"mq"`
	GitLab   GitLab `json:"gitlab"`
}

type PGSQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	DBName   string `json:"db_name"`
	Password string `json:"password"`
}

type MQ struct {
	Address  string `json:"address"`
	MaxRetry int    `json:"max_retry"`
	Topics   `json:"topics"`
}

type GitLab struct {
	RootToken    string        `json:"root_token"`
	Endponit     string        `json:"endpoint"`
	CountPerPage int           `json:"count_per_page"`
	RefreshTime  time.Duration `json:"refresh_time"`
}

type Topics struct {
	Statistics string `json:"statistics" required:"true"`
	GitLab     string `json:"gitlab" required:"true"`
}

func (cfg *Config) GetMQConfig() mq.MQConfig {
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
