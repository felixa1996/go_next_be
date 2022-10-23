package config

import "github.com/spf13/viper"

type Config struct {
	// App
	AppName string `mapstructure:"APP_NAME"`
	Port    string `mapstructure:"PORT"`
	Timeout int    `mapstructure:"TIMEOUT"`
	Debug   bool   `mapstructure:"DEBUG"`
	// Keycloak
	KeycloakHost         string `mapstructure:"KEYCLOAK_HOST"`
	KeycloakRealm        string `mapstructure:"KEYCLOAK_REALM"`
	KeycloakClientId     string `mapstructure:"KEYCLOAK_CLIENT_ID"`
	KeycloakClientSecret string `mapstructure:"KEYCLOAK_CLIENT_SECRET"`
	// Database
	MongoUri string `mapstructure:"MONGODB_URI"`
	MongoDB  string `mapstructure:"MONGODB_DB"`
	// Sqs
	// Upsert Company Queue
	SqsCompanyUpsertUrl              string `mapstructure:"SQS_COMPANY_UPSERT_URL"`
	SqsCompanyUpsertMaxNumberMessage int64  `mapstructure:"SQS_COMPANY_UPSERT_MAX_NUMBER_MESSAGE"`
	SqsCompanyWaitTimeOutSeconds     int64  `mapstructure:"SQS_COMPANY_UPSERT_WAIT_TIMEOUT"`
}

func LoadConfigFromEnv() (config Config) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return
}