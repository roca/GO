package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
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
	States           map[string]bool
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

	a := app{
		States: make(map[string]bool),
	}

	http.HandleFunc("/", a.index)
	http.HandleFunc("/callback", a.callback)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("ListenAndServe error: %s\n", err)
	}
}

func (a *app) index(w http.ResponseWriter, r *http.Request) {

	url, state, err := GetAuthorizationURL()
	if err != nil {
		returnError(w, fmt.Errorf("SetAuthorizationURL error %s", err))
		return
	}
	a.AuthorizationURL = url
	a.States[state] = true

	err = templates["index.html"].Execute(w, a)
	if err != nil {
		returnError(w, fmt.Errorf("Template Execute error %s", err))
		return
	}

}

func (a *app) callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if _, ok := a.States[state]; !ok {
		returnError(w, fmt.Errorf("Invalid state"))
		return
	}

	configFileBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		returnError(w, fmt.Errorf("Error reading config file %s", err))
		return
	}

	config := ReadConfig(configFileBytes)
	oidcEndpoint := config.Url
	discovery, err := oidc.ParseDiscovery(oidcEndpoint + "/.well-known/openid-configuration")
	if err != nil {
		returnError(w, fmt.Errorf("ParseDiscovery error: %s", err))
		return
	}

	delete(a.States, state)

	token, _, err := getTokenFromCode(
		discovery.TokenEndpoint,
		discovery.JwksURI,
		config.Apps["app1"].RedirectURIs[0],
		config.Apps["app1"].ClientID,
		config.Apps["app1"].ClientSecret,
		code,
	)
	if err != nil {
		returnError(w, fmt.Errorf("getTokenFromCode error: %s", err))
		return
	}

	req, err := http.NewRequest("GET", discovery.UserinfoEndpoint, nil)
	if err != nil {
		returnError(w, fmt.Errorf("NewRequest error: %s", err))
		return
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		returnError(w, fmt.Errorf("Do request error: %s", err))
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		returnError(w, fmt.Errorf("ReadAll error: %s", err))
		return
	}

	// var user users.User

	// err = yaml.Unmarshal(body, &user)
	// if err != nil {
	// 	returnError(w, fmt.Errorf("Unmarshal error: %s", err))
	// 	return
	// }

	// out, err = json.Marshal(user)
	// if err != nil {
	// 	returnError(w, fmt.Errorf("Marshal error: %s", err))
	// 	return
	// }

	w.Write([]byte(fmt.Sprintf("Token received. User info: %s", body)))
}

func returnError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
	fmt.Printf("Error: %s\n", err)
}

func GetAuthorizationURL() (string, string, error) {

	err := LoadTemplates()
	if err != nil {
		return "", "", err
	}

	configFileBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return "", "", err
	}

	config := ReadConfig(configFileBytes)
	oidcEndpoint := config.Url
	discovery, err := oidc.ParseDiscovery(oidcEndpoint + "/.well-known/openid-configuration")
	if err != nil {
		return "", "", err
	}

	state, err := oidc.GetRandomString(64)
	if err != nil {
		return "", "", err
	}

	authorizationURL := fmt.Sprintf("%s?client_id=%s&response_type=code&scope=openid&redirect_uri=%s&state=%s",
		discovery.AuthorizationEndpoint,
		config.Apps["app1"].ClientID,
		config.Apps["app1"].RedirectURIs[0],
		state,
	)

	return authorizationURL, state, nil
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
