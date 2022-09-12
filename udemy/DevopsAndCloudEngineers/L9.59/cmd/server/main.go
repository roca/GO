package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/ssh"
)

func main() {
	var (
		err error
	)
	authorizedKeysBytes, err := ioutil.ReadFile("mykey.pub")
	if err != nil {
		log.Fatalf("Failed to load authorized_keys, err: %v", err)
	}
	privateKey, err := ioutil.ReadFile("server.pem")
	if err != nil {
		log.Fatalf("Failed to load private key, err: %v", err)
	}

	if err = ssh.StartServer(privateKey, authorizedKeysBytes); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

}