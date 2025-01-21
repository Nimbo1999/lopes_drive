package config

import (
	"log"

	"github.com/spf13/viper"
)

func ReadEnvs() {
	viper.SetConfigFile(".env")
	viper.SetDefault("PORT", "80")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

type EnvironmentVariables struct {
	AwsAccessKey string
	AwsSecretKey string
	Port         string
}

func CreateEnvironment() *EnvironmentVariables {
	return &EnvironmentVariables{
		AwsAccessKey: viper.GetString("AWS_ACCESS_KEY_ID"),
		AwsSecretKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
		Port:         viper.GetString("PORT"),
	}
}
