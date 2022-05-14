package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Port   int    `toml:"port"`
		JWTKey string `toml:"jwt_key"`
	} `toml:"app"`
	Database struct {
		Driver string `toml:"driver"`
		Db_url string `toml:"db_url"`
	} `toml:"database"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.App.Port = 8081

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("No config file found, using default config")

		return &defaultConfig
	}
	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("Failed to parse config file, using default config")
		return &defaultConfig
	}
	return &finalConfig
}
