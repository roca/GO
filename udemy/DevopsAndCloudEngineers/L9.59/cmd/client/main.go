package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	scp "github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

func main() {
	var (
		cmd    *string
		source *string
		dist   *string
	)

	cmd = flag.String("c", "", "command to run")
	source = flag.String("s", "", "source file")
	dist = flag.String("d", "", "target destination file")

	flag.Parse()

	if *cmd == "" {
		flag.Usage()
		os.Exit(1)
	}

	var (
		err error
	)
	privateKey, err := ioutil.ReadFile("mykey.pem")
	if err != nil {
		log.Fatalf("Failed to load mykey.pem, err: %v", err)
	}
	publicKey, err := ioutil.ReadFile("server.pub")
	if err != nil {
		log.Fatalf("Failed to server.pub, err: %v", err)
	}

	privateKeyParsed, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse mykey.pem, err: %v", err)
	}

	publicKeyParsed, _, _, _, err := ssh.ParseAuthorizedKey(publicKey)
	if err != nil {
		log.Fatalf("Failed to parse server.pub, err: %v", err)
	}

	clientConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKeyParsed),
		},
		HostKeyCallback: ssh.FixedHostKey(publicKeyParsed),
	}

	client, err := ssh.Dial("tcp", "localhost:2022", clientConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()
	if *cmd != "" {
		switch *cmd {
		case "whoami":
			out, err := session.Output("whoami")
			if err != nil {
				log.Fatalf("Session output error: %s", err)
			}
			log.Printf("Output is: %s", string(out))

		case "copy":
			fmt.Printf("source: %s to destination: %s\n", *source, *dist)
			SendFile(client, *source , *dist)
		default:
			log.Fatalf("Can't handle this command: %s", *cmd)
		}
	}
}

func SendFile(client *ssh.Client, source, dist string) error {
	scpClient, err := scp.NewClientBySSH(client)
	defer scpClient.Close()
	if err != nil {
		fmt.Println("Error creating new SSH session from existing connection", err)
	}
	f, _ := os.Open(source)
	defer f.Close()
	err = scpClient.Connect()
	if err != nil {
		return err
	}
	err = scpClient.CopyFromFile(context.Background(), *f, dist, "0655")
	if err != nil {
		return err
	}
	return nil
}
