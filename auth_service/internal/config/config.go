package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env           string        `yaml:"env" env-default:"local"`
	DataProvider  string        `yaml:"data_provider" env-required:"true"`
	GRPC          GRPCConfig    `yaml:"grpc"`
	MySQLSettings MySQLConfig   `yaml:"mysql_settings"`
	TokenTTL      time.Duration `yaml:"token_ttl"`
}

type GRPCConfig struct {
	Port int `yaml:"port"`
}

// region databases providers

type MySQLConfig struct {
	Address  string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"db"`
}

type StorageConfig struct {
	DataProvider string

	// для других бд можно добавить другие настройки
	MySQLSettings MySQLConfig
}

// endregion

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
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
