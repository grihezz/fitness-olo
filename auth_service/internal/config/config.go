// Package config provides functionality for loading application configuration.
// Пакет config предоставляет функциональность для загрузки конфигурации приложения.
package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

// Config represents the application configuration.
type Config struct {
	Env           string        `yaml:"env" env-default:"local"`
	DataProvider  string        `yaml:"data_provider" env-required:"true"`
	GRPC          GRPCConfig    `yaml:"grpc"`
	MySQLSettings MySQLConfig   `yaml:"mysql_settings"`
	TokenTTL      time.Duration `yaml:"token_ttl"`
}

// GRPCConfig represents gRPC configuration.
type GRPCConfig struct {
	Port int `yaml:"port"`
}

// MySQLConfig represents MySQL database configuration.
type MySQLConfig struct {
	Address  string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"db"`
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

// fetchConfigPath fetches the path to the configuration file.
func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	if res == "" {
		res = "auth_service/resourse/local/config.yml"
	}

	return res
}
