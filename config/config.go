package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init (env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("default")
	config.AddConfigPath("config/")
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("can't read config %s", env))
	}

	// TODO: add env
}

func Config() *viper.Viper {
	return config
}
