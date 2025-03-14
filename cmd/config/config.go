package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yml:"env env-default:"local"`
	HTTPServer `yml:"http_server"`
}

type HTTPServer struct {
	Address string `yml:"address" env-default:"localhost:8080"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_FILE_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file is not exist: %s", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("can not read config: %s", err)
	}

	return &cfg
}
