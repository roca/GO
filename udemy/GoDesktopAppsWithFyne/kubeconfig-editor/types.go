package main

type ClusterConFig struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data"`
	Server                   string `yaml:"server"`
}

type Cluster struct {
	Cluster ClusterConFig `yaml:"cluster"`
	Name    string        `yaml:"name"`
}

type ContextBinding struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

type Context struct {
	Context ContextBinding `yaml:"context"`
	Name    string         `yaml:"name"`
}

type AuthProviderConfig struct {
	AccessToken string `yaml:"access-token"`
	CmdArgs     string `yaml:"cmd-args"`
	Expiry      string `yaml:"expiry"`
	ExpiryKey   string `yaml:"expiry-key"`
	TokenKey    string `yaml:"token-key"`
}

type AuthProvider struct {
	Config AuthProviderConfig `yaml:"config"`
	Name   string             `yaml:"name"`
}

type UserConfig struct {
	ClientCertificateData string       `yaml:"client-certificate-data"`
	ClientKeyData         string       `yaml:"client-key-data"`
	AuthProvider          AuthProvider `yaml:"auth-provider"`
}

type User struct {
	User UserConfig `yaml:"user"`
	Name string     `yaml:"name"`
}

type kubeconfig struct {
	ApiVersion     string    `yaml:"apiVersion"`
	Clusters       []Cluster `yaml:"clusters"`
	Contexts       []Context `yaml:"contexts"`
	CurrentContext string    `yaml:"current-context"`
	Kind           string    `yaml:"kind"`
	Preferences    struct{}  `yaml:"preferences"`
	Users          []User    `yaml:"users"`
}
