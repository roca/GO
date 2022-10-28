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
