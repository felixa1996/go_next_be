package iam

import (
	"context"

	"github.com/Nerzal/gocloak/v11"
	"github.com/spf13/viper"
)

type KeycloakIAM struct {
	Client       gocloak.GoCloak
	host         string
	clientId     string
	clientSecret string
	realm        string
}

func NewKeycloakIAM() KeycloakIAM {
	host := viper.GetString("HOST")

	client := gocloak.NewClient(host)

	return KeycloakIAM{
		Client:       client,
		host:         host,
		clientId:     viper.GetString("CLIENT_ID"),
		clientSecret: viper.GetString("CLIENT_SECRET"),
		realm:        viper.GetString("REALM"),
	}
}

func (k *KeycloakIAM) RetrospectToken(context context.Context, accessToken string) (*gocloak.RetrospecTokenResult, error) {
	return k.Client.RetrospectToken(context, accessToken, k.clientId, k.clientSecret, k.realm)
}
