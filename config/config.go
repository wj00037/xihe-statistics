package config

import (
	"errors"
	"os"
	"regexp"
	"strings"
	"time"

	"sigs.k8s.io/yaml"

	"github.com/opensourceways/kafka-lib/agent"
)

var reIpPort = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}:[1-9][0-9]*$`)

func LoadFromYaml(path string, cfg interface{}) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, cfg)
}

func LoadConfig(path string, cfg interface{}) error {
	if err := LoadFromYaml(path, cfg); err != nil {
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
	MQTopics Topics `json:"mq_topics"    required:"true"`

	GitLab GitLab `json:"gitlab"`
}

type PGSQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	DBName   string `json:"db_name"`
	Password string `json:"password"`
	DBCert   string `json:"db_cert"`
}

type MQ struct {
	Address   string `json:"address"   required:"true"`
	Version   string `json:"version"`
	MQCert    string `json:"mq_cert"   required:"true"`
	Username  string `json:"user_name" required:"true"`
	Password  string `json:"password"  required:"true"`
	Algorithm string `json:"algorithm" required:"true"`
}

type GitLab struct {
	RootToken    string        `json:"root_token"`
	Endponit     string        `json:"endpoint"`
	CountPerPage int           `json:"count_per_page"`
	RefreshTime  time.Duration `json:"refresh_time"`
}

type Topics struct {
	Statistics      string `json:"statistics"          required:"true"`
	GitLab          string `json:"gitlab"              required:"true"`
	Training        string `json:"training"            required:"true"`
	Cloud           string `json:"cloud"               required:"true"`
	BigModelStarted string `json:"bigmodel_started"    required:"true"`
}

func (cfg *Config) GetKfkConfig() agent.Config {
	return agent.Config{
		Address:   cfg.MQ.Address,
		Version:   cfg.MQ.Version,
		MQCert:    cfg.MQ.MQCert,
		Username:  cfg.MQ.Username,
		Password:  cfg.MQ.Password,
		Algorithm: cfg.MQ.Algorithm,
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
