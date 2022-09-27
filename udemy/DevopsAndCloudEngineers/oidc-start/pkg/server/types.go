package server

import (
	"time"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/users"
)

type Config struct {
	Apps      map[string]AppConfig `yaml:"apps"`
	Url       string               `yaml:"url"`
	LoadError error
}
type AppConfig struct {
	ClientID     string   `yaml:"clientID"`
	ClientSecret string   `yaml:"clientSecret"`
	Issuer       string   `yaml:"issuer"`
	RedirectURIs []string `yaml:"redirectURIs"`
}

type LoginRequest struct {
	ClientID     string
	RedirectURI  string
	Scope        string
	ResponseType string
	State        string
	CodeIssuedAt time.Time
	User         users.User
	AppConfig    AppConfig
}
