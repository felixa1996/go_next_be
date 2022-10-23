package iam

import (
	"context"

	"github.com/Nerzal/gocloak/v11"
	"github.com/felixa1996/go_next_be/app/config"
)

type KeycloakIAM struct {
	Client       gocloak.GoCloak
	host         string
	clientId     string
	clientSecret string
	realm        string
}

func NewKeycloakIAM(config config.Config) KeycloakIAM {
	client := gocloak.NewClient(config.KeycloakHost)

	return KeycloakIAM{
		Client:       client,
		host:         config.KeycloakHost,
		clientId:     config.KeycloakClientId,
		clientSecret: config.KeycloakClientSecret,
		realm:        config.KeycloakRealm,
	}
}

func (k *KeycloakIAM) RetrospectToken(context context.Context, accessToken string) (*gocloak.RetrospecTokenResult, error) {
	return k.Client.RetrospectToken(context, accessToken, k.clientId, k.clientSecret, k.realm)
}
