package main

import (
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/wardviaene/golang-for-devops-course/ssh-demo"
)

func main() {
	var (
		token  string
		pubKey string
		err    error
	)

	if pubKey, err = generateKey(); err != nil {
		fmt.Printf("generateKey() error: %s\n", err)
		os.Exit(1)
	}
	if token, err = getToken(); err != nil {
		fmt.Printf("getToken() error: %s\n", err)
		os.Exit(1)
	}
	if err = launchInstance(token, pubKey); err != nil {
		fmt.Printf("launchInstance() error: %s\n", err)
		os.Exit(1)
	}
}

func generateKey() (string, error) {
	var (
		privateKey []byte
		publicKey  []byte
		err        error
	)

	if privateKey, publicKey, err = ssh.GenerateKeys(); err != nil {
		return "", err
	}

	if err = os.WriteFile("myKey.pem", privateKey, 0600); err != nil {
		return "", err
	}

	if err = os.WriteFile("myKey.pub", publicKey, 0644); err != nil {
		return "", err
	}

	return string(publicKey), nil
}

func getToken() (string, error) {
	_, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		return "", err
	}
	return "", nil
}

func launchInstance(token, pubKey string) error {
	return nil
}
