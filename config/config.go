package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init (env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("config/")
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("can't read config %s; %v", env, err))
	}

	// TODO: add env
}

func Config() *viper.Viper {
	return config
}
