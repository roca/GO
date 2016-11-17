package data

import (
	"crypto/tls"
	"flag"
	"net/http"
)

var dataServiceUrl = flag.String("dataservice", "https://localhost:4000", "Address of the data service provider")

func init() {
	tr := http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	http.DefaultClient = &http.Client{Transport: &tr}
}
