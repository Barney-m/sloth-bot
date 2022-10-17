package config

import (
	"github.com/spf13/viper"
)

type SlothConfig struct {
	BotToken  string `mapstructure:"BOT_TOKEN"`
	BotPrefix string `mapstructure:"BOT_PREFIX"`
}

var Config *SlothConfig

func LoadConfig() (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("sloth")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&Config)
	return
}
