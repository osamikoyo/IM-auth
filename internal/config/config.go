package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	Port uint `yaml:"port"`
	Hostname string `yaml:"hostname"`
	DSN string `yaml:"dsn"`
	JwtKey string `yaml:"jwt_key"`
	RpcQueName string `yaml:"rpc_que_name"`
	AmqpConnect string `yaml:"amqp_connect"`
}

func Load(cfgPath string) (*Config, error) {
	var cfg Config

	file, err := os.Open(cfgPath)
	if err != nil{
		return nil, err
	}

	body, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(body, &cfg)
	return &cfg, err
}