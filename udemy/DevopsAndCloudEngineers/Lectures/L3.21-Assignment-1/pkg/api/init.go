package api

import (
	"io"
	"net/http"
)

type ClientInterface interface {
	Get(url string) (*http.Response, error)
	Post(url string, contenType string, body io.Reader) (resp *http.Response, err error)
}

type Options struct {
	Password string
	LoginURL string
}

type APIIface interface {
	DoGetRequest(requestURL string) (IResponse, error)
}

type api struct {
	Options Options
	Client  ClientInterface
}

func New(options Options) APIIface {
	return api{
		Options: options,
		Client: &http.Client{
			Transport: &MyJWTTransport{
				transport:  http.DefaultTransport,
				password:   options.Password,
				loginURL:   options.LoginURL,
				HTTPClient: &http.Client{},
			},
		},
	}
}
