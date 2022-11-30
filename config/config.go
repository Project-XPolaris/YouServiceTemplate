package config

import (
	"github.com/allentom/harukap/config"
	"github.com/sirupsen/logrus"
	"os"
)

var DefaultConfigProvider *config.Provider

func InitConfigProvider() error {
	var err error
	customConfigPath := os.Getenv("CONFIG_PATH")
	if customConfigPath != "" {
		logrus.Info("using custom config path: ", customConfigPath)
	}
	DefaultConfigProvider, err = config.NewProvider(func(provider *config.Provider) {
		ReadConfig(provider)
	}, customConfigPath)
	return err
}

var Instance Config

type Config struct {
}

func ReadConfig(provider *config.Provider) {
	configer := provider.Manager
	configer.SetDefault("addr", ":8000")
	configer.SetDefault("application", "My Service")
	configer.SetDefault("instance", "main")

	Instance = Config{}
}
