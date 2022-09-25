package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"net/http"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/oidc-start/pkg/oidc"
	"gopkg.in/yaml.v2"
)

var (
	//go:embed templates/*
	files     embed.FS
	templates map[string]*template.Template
)

const (
	templatesDir = "templates"
)

type app struct {
	AuthorizationURL string
}

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

func main() {

	a := app{}

	http.HandleFunc("/", a.index)
	http.HandleFunc("/callback", a.callback)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("ListenAndServe error: %s\n", err)
	}
}

func (a *app) index(w http.ResponseWriter, r *http.Request) {

	url, err := GetAuthorizationURL()
	if err != nil {
		returnError(w, fmt.Errorf("SetAuthorizationURL error %s", err))
		return
	}
	a.AuthorizationURL = url

	err = templates["index.html"].Execute(w, a)
	if err != nil {
		returnError(w, fmt.Errorf("Template Execute error %s", err))
		return
	}

}

func (a *app) callback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("callback. Code is : " + r.URL.Query().Get("code")))
}

func returnError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
	fmt.Printf("Error: %s\n", err)
}

func GetAuthorizationURL() (string, error) {

	err := LoadTemplates()
	if err != nil {
		return "", err
	}

	configFileBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return "", err
	}

	config := ReadConfig(configFileBytes)
	oidcEndpoint := config.Url
	discovery, err := oidc.ParseDiscovery(oidcEndpoint + "/.well-known/openid-configuration")
	if err != nil {
		return "", err
	}

	state, err := oidc.GetRandomString(64)
	if err != nil {
		return "", err
	}

	authorizationURL := fmt.Sprintf("%s?client_id=%s&response_type=code&scope=openid&redirect_uri=%s&state=%s",
		discovery.AuthorizationEndpoint,
		config.Apps["app1"].ClientID,
		config.Apps["app1"].RedirectURIs[0],
		state,
	)

	return authorizationURL, nil
}

func LoadTemplates() error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(files, templatesDir+"/"+tmpl.Name())
		if err != nil {
			return err
		}

		templates[tmpl.Name()] = pt
	}
	return nil
}

func ReadConfig(bytes []byte) Config {
	var config Config
	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		config.LoadError = err
	}
	return config
}
