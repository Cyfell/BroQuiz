package config

import (
	"github.com/Cyfell/BroQuiz/internal/api"
	"github.com/spf13/viper"
)

type Config struct {
	API api.Config `mapstructure:"api"`
}

func ReadConfig(filename string) (Config, error) {
	var config Config

	viper.SetConfigName(filename)
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/broquiz/")
	viper.AddConfigPath("../../configs/")
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
