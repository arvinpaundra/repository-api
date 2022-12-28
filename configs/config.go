package configs

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("err while reading config file: %v", err)
	}

	return viper.GetString(key)
}
