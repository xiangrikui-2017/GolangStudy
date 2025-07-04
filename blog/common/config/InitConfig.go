package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	DB  DBConfig
	Log LogConfig
	Jwt JwtConfig
}

var Conf = &Config{}

func InitConfig(configPath string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	})

	return Conf
}
