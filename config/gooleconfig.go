package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupGoogleConfig() *oauth2.Config{
	conf := &oauth2.Config{
		ClientID: "",
		ClientSecret: "",
		RedirectURL: "http://localhost:8080/google/callback",
		Scopes: []string{
			"",
			"",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}