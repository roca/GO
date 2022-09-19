package server

type Config struct {
	Apps      map[string]AppConfig `yaml:"apps"`
	Url       string               `yaml:"url"`
	LoadError error
}
type AppConfig struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	Issuer       string   `yaml:"issuer"`
	RedirectURIs []string `yaml:"redirect_uris"`
}
