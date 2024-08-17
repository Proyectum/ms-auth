package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName("application.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
