package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "it's working")
}

func showCommonName(w http.ResponseWriter, req *http.Request) {
	var commonName string
	if req.TLS != nil && len(req.TLS.VerifiedChains) > 0 && len(req.TLS.VerifiedChains[0]) > 0 {
		commonName = req.TLS.VerifiedChains[0][0].Subject.CommonName
	}
	fmt.Fprintf(w, "Your common Name: %s", commonName)
}

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

	http.HandleFunc("/", index)
	http.HandleFunc("/common-name", showCommonName)
	server := http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  ca,
			MinVersion: tls.VersionTLS13,
		},
	}
	err = server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal("ListenAndServeTLS error: ", err)
	}
}
