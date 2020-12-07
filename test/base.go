package test

import "github.com/spf13/viper"

func setConfig() {
	viper.SetConfigName("test")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../configs/environment")
	viper.ReadInConfig()
}
