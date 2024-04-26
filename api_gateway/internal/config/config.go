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

type OloService struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// ToStr returns a string representation of the HTTPConfig, which consists of host and port concatenated with a colon.
func (h HTTPConfig) ToStr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

type Config struct {
	Env         string      `yaml:"env" env-default:"local"`
	AuthService AuthService `yaml:"auth_service"`
	OloService  OloService  `yaml:"olo_service"`
	HTTP        HTTPConfig  `yaml:"http"`
}

// MustLoad loads the configuration from the specified path.
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file is not exist " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config " + err.Error())
	}
	return &cfg
}

// fetchConfigPath fetches the configuration file path.
func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "api_gateway/resourse/config.yml")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	if res == "" {
		res = "api_gateway/resourse/local/config.yml"
	}
	return res
}
