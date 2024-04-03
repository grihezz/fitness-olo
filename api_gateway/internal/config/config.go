package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type AuthService struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (h HTTPConfig) ToStr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

type Config struct {
	Env         string      `yaml:"env" env-default:"local"`
	AuthService AuthService `yaml:"auth_service"`
	HTTP        HTTPConfig  `yaml:"http"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file is not exist" + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "api_gateway/resourse/config.yml")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	if res == "" {
		res = "api_gateway/resourse/dev/config.yml"
	}
	return res
}
