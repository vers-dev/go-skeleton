package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type (
	Config struct {
		Server *Server
		Db     *Db
	}

	Server struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {

	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	fmt.Println("config Successful load ")

	return configInstance
}
