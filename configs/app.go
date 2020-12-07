package configs

import (
	"os"

	"github.com/spf13/viper"
)

var runtimeEnv string

func Init() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
}

func IsLocal() bool {
	return runtimeEnv == "local"
}

func loadConfig() error {
	runtimeEnv = os.Getenv("APP_ENV")

	filename := runtimeEnv
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/environment")
	err := viper.ReadInConfig()
	return err
}
