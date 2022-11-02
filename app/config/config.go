package config

import "github.com/spf13/viper"

type Config struct {
	// App
	AppName   string `mapstructure:"APP_NAME"`
	Port      string `mapstructure:"PORT"`
	Timeout   int    `mapstructure:"TIMEOUT"`
	Debug     bool   `mapstructure:"DEBUG"`
	TraceType string `mapstructure:"TRACE_TYPE"`
	// Keycloak
	KeycloakHost         string `mapstructure:"KEYCLOAK_HOST"`
	KeycloakRealm        string `mapstructure:"KEYCLOAK_REALM"`
	KeycloakClientId     string `mapstructure:"KEYCLOAK_CLIENT_ID"`
	KeycloakClientSecret string `mapstructure:"KEYCLOAK_CLIENT_SECRET"`
	// Database
	MongoUri string `mapstructure:"MONGODB_URI"`
	MongoDB  string `mapstructure:"MONGODB_DB"`
	// Newrelic
	NewrelicLicenseKey string `mapstructure:"NEWRELIC_LICENSEKEY"`
	// Minio
	MinioEndpoint        string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey       string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretAccessKey string `mapstructure:"MINIO_SECRET_ACCESS_KEY"`
	MinioSSL             bool   `mapstructure:"MINIO_SSL"`
	MinioBaseUrl         string `mapstructure:"MINIO_BASE_URL"`
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
