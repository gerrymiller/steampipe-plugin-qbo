package qbo

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type qboConfig struct {
	BaseURL *string `cty:"baseURL"`
	ClientId *string `cty:"clientId"`
    ClientSecret *string `cty:"clientSecret"`
    RealmId *string `cty:"realmId"`
	AccessToken *string `cty:"accessToken"`
	RefreshToken *string `cty:"refreshToken"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"baseURL": {
		Type: schema.TypeString,
	},
	"clientId": {
		Type: schema.TypeString,
	},
	"clientSecret": {
		Type: schema.TypeString,
	},
	"realmId": {
		Type: schema.TypeString,
	},
	"accessToken": {
		Type: schema.TypeString,
	},
	"refreshToken": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &qboConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) qboConfig {
	if connection == nil || connection.Config == nil {
		return qboConfig{}
	}
	config, _ := connection.Config.(qboConfig)
	return config
}