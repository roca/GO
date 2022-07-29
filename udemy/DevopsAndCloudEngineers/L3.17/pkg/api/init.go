package api

import "net/http"

type Options struct {
	Password string
	LoginURL string
}

type APIIface interface {
	DoGetRequest(requestURL string) (IResponse, error)
}

type api struct {
	Optiomns Options
	Client   http.Client
}

func New(options Options) APIIface {
	return api{
		Optiomns: options,
		Client: http.Client{
			Transport: &MyJWTTransport{
				transport: http.DefaultTransport,
				password:  options.Password,
				loginURL:  options.LoginURL,
			},
		},
	}
}
