package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func loadConfig() {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "standalone"
	}
	viper.SetConfigName(fmt.Sprintf("application-%s.yaml", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
