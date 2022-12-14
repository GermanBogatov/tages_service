package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	IsDebug       bool `yaml:"is-debug" env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `yaml:"is-development" env:"IS_DEV" env-default:"false"`
	GRPC          struct {
		IP   string `yaml:"ip" env:"GRPC-IP"`
		Port int    `yaml:"port" env:"GRPC-PORT"`
	} `yaml:"grpc"`
	PostgreSQL struct {
		Username string `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
		Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
	} `yaml:"postgresql"`
	AppConfig struct {
		LogLevel string `yaml:"log-level" env:"LOG_LEVEL" env-default:"trace"`
	} `yaml:"app"`
}

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(&configPath, FlagConfigPathName, "configs/config.yaml", "this is app config file")
		flag.Parse()

		log.Print("config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "German Bogatov - Tages Service"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
