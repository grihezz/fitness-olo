package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env           string      `yaml:"env" env-default:"local"`
	GRPC          GRPCConfig  `yaml:"grpc"`
	MySQLSettings MySQLConfig `yaml:"mysql_settings"`
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

// endregion

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
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	if res == "" {
		res = "olo_service/resourse/local/config.yml"
	}
	return res
}
