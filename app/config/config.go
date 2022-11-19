package config

import "github.com/spf13/viper"

func LoadConfig() (config Config) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetString("ENV_TYPE") == "vault" {
		config = LoadVaultConfigFromEnv()
		return config
	}

	return LoadConfigFromEnv()
}
