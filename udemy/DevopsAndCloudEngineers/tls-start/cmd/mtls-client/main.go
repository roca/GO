package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	caBytes, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("Error reading ca cert: %s\n", err)
		return
	}

	ca := x509.NewCertPool()
	if ok := ca.AppendCertsFromPEM(caBytes); !ok {
		log.Fatalf("Error appending ca cert: %s\n", err)
	}
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatalf("Error loading client cert: %s\n", err)
	}

	client := http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      ca,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	resp, err := client.Get("https://go-demo.localtest.me:443/common-name")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Body (status %d): %s\n", resp.StatusCode, body)
}
