package common

import "github.com/spf13/viper"

func ReadConfigFromEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
