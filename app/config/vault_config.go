package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
)

type VaultConfig struct {
	// App
	Debug     bool   `mapstructure:"DEBUG"`
	EnvType   string `mapstructure:"ENV_TYPE"`
	Address   string `mapstructure:"VAULT_ADDRESS"`
	RoleId    string `mapstructure:"VAULT_ROLE_ID"`
	SecretId  string `mapstructure:"VAULT_SECRET_ID"`
	SecretUrl string `mapstructure:"VAULT_SECRET_URL"`
}

func LoadVaultConfigFromEnv() (config Config) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var vaultConfigEnv *VaultConfig
	err = viper.Unmarshal(&vaultConfigEnv)
	if err != nil {
		panic(err)
	}

	configVault := &vault.Config{
		Address: vaultConfigEnv.Address,
	}

	client, err := vault.NewClient(configVault)
	if err != nil {
		panic(fmt.Errorf("Unable to initialize Vault client: %w", err))
	}

	appRoleAuth, err := auth.NewAppRoleAuth(
		vaultConfigEnv.RoleId,
		&auth.SecretID{FromString: vaultConfigEnv.SecretId},
		// auth.WithWrappingToken(), // Only required if the secret ID is response-wrapped.
	)
	if err != nil {
		panic(fmt.Errorf("unable to initialize AppRole auth method: %w", err))
	}

	authInfo, err := client.Auth().Login(context.Background(), appRoleAuth)
	if err != nil {
		panic(fmt.Errorf("unable to login to AppRole auth method: %w", err))
	}
	if authInfo == nil {
		panic(fmt.Errorf("no auth info was returned after login"))
	}

	secret, err := client.KVv2("kv").Get(context.Background(), vaultConfigEnv.SecretUrl)
	if err != nil {
		panic(fmt.Errorf("unable to read secret: %w", err))
	}

	return SetVaultToEnv(secret)
}

func SetVaultToEnv(secret *vault.KVSecret) (config Config) {
	// App
	config.AppName = GetStringFromSecret(secret, "APP_NAME")
	config.Port = GetStringFromSecret(secret, "PORT")
	config.Timeout = GetIntFromSecret(secret, "TIMEOUT")
	config.TraceType = GetStringFromSecret(secret, "TRACE_TYPE")

	// Mongo
	config.MongoUri = GetStringFromSecret(secret, "MONGODB_URI")
	config.MongoDB = GetStringFromSecret(secret, "MONGODB_DB")

	// KeyCloakIAM
	config.KeycloakClientId = GetStringFromSecret(secret, "KEYCLOAK_CLIENT_ID")
	config.KeycloakClientSecret = GetStringFromSecret(secret, "KEYCLOAK_CLIENT_SECRET")
	config.KeycloakRealm = GetStringFromSecret(secret, "KEYCLOAK_REALM")
	config.KeycloakHost = GetStringFromSecret(secret, "KEYCLOAK_REALM")

	// Aws
	// For Aws,it's read conf from os env because of the security options
	os.Setenv("AWS_ACCESS_KEY_ID", GetStringFromSecret(secret, "AWS_ACCESS_KEY_ID"))
	os.Setenv("AWS_SECRET_ACCESS_KEY", GetStringFromSecret(secret, "AWS_SECRET_ACCESS_KEY"))

	// Newrelic
	config.NewrelicLicenseKey = GetStringFromSecret(secret, "NEWRELIC_LICENSEKEY")

	// Minio
	config.MinioEndpoint = GetStringFromSecret(secret, "MINIO_ENDPOINT")
	config.MinioAccessKey = GetStringFromSecret(secret, "MINIO_ACCESS_KEY")
	config.MinioSecretAccessKey = GetStringFromSecret(secret, "MINIO_SECRET_ACCESS_KEY")
	config.MinioSSL = GetBoolFromSecret(secret, "MINIO_SSL")
	config.MinioBaseUrl = GetStringFromSecret(secret, "MINIO_BASE_URL")

	// Sqs
	config.SqsCompanyUpsertUrl = GetStringFromSecret(secret, "SQS_COMPANY_UPSERT_URL")
	config.SqsCompanyUpsertMaxNumberMessage = GetInt64FromSecret(secret, "SQS_COMPANY_UPSERT_MAX_NUMBER_MESSAGE")
	config.SqsCompanyWaitTimeOutSeconds = GetInt64FromSecret(secret, "SQS_COMPANY_UPSERT_WAIT_TIMEOUT")

	return config
}

func GetStringFromSecret(secret *vault.KVSecret, key string) string {
	value, ok := secret.Data[key].(string)
	if !ok {
		panic(fmt.Errorf("Failed to get string from secret: %#v", key))
	}
	return value
}

func GetInt64FromSecret(secret *vault.KVSecret, key string) int64 {
	value, ok := secret.Data[key].(string)
	if !ok {
		panic(fmt.Errorf("Failed to get int64 from secret: %#v", key))
	}
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Errorf("Failed to convert from string to int64 from secret: %#v", key))
	}
	return valueInt64
}

func GetIntFromSecret(secret *vault.KVSecret, key string) int {
	value, ok := secret.Data[key].(string)
	if !ok {
		panic(fmt.Errorf("Failed to get int64 from secret: %#v", key))
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Errorf("Failed to convert from string to int64 from secret: %#v", key))
	}
	return valueInt
}

func GetBoolFromSecret(secret *vault.KVSecret, key string) bool {
	value, ok := secret.Data[key].(string)
	if !ok {
		panic(fmt.Errorf("Failed to get bool from secret: %#v", key))
	}
	return value == "true"
}
