package server

import "testing"

//  var configYamlFile = `
// apps:
// 	- name: app1
// 		client_id: client01
// 		client_secret: clientsecret
// 		issuer: https://localhost:8080
// 		redirect_urls:
// 			- https://localhost:8080/callback
// 	url: https://localhost:8080
//  `

 var configYamlFile = `
url: https://localhost:8080
apps:
  app1:
    clientID: ClientID
    clientSecret: ClientSecret
    issuer: https://localhost:8080
    redirectURIs:
      - https://localhost:8080/callback
 `

func TestReadConfig(t *testing.T) {
	config := ReadConfig([]byte(configYamlFile))

	if config.LoadError != nil {
		t.Errorf("Expected no error, got %v", config.LoadError)
	}
	if config.Url != "https://localhost:8080" {
		t.Errorf("Expected https://localhost:8080, got %v", config.Url)
	}
	if len(config.Apps) != 1 {
		t.Errorf("Expected 1 app, got %v", len(config.Apps))
	}
	if config.Apps["app1"].ClientID != "ClientID" {
		t.Errorf("Expected client_id, got %v", config.Apps["app1"].ClientID)
	}
	if config.Apps["app1"].ClientSecret != "ClientSecret" {
		t.Errorf("Expected client_secret, got %v", config.Apps["app1"].ClientSecret)
	}
	if config.Apps["app1"].Issuer != "https://localhost:8080" {
		t.Errorf("Expected https://localhost:8080, got %v", config.Apps["app1"].Issuer)
	}
	if len(config.Apps["app1"].RedirectURIs) != 1 {
		t.Errorf("Expected 1 redirect uri, got %v", len(config.Apps["app1"].RedirectURIs))
	}
	if config.Apps["app1"].RedirectURIs[0] != "https://localhost:8080/callback" {
		t.Errorf("Expected https://localhost:8080/callback, got %v", config.Apps["app1"].RedirectURIs[0])
	}
}
