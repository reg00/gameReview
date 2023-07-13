package config

import (
	"github.com/spf13/viper"
)

var gConfig Configuration

func Register() *Configuration {
	return &gConfig
}

func LoadConfig() (Configuration, error) {
	viper := viper.New()
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".\\internal\\infrastructure\\config")
	viper.AddConfigPath("internal/infrastructure/config")
	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, err
	}

	var cfg Configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return Configuration{}, err
	}

	gConfig = cfg
	return gConfig, nil
}
