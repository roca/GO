package utils

import "net/http"

// GetResponseHeaders return default response headers
func GetResponseHeaders() http.Header {

	headers := http.Header{}

	headers.Add("key string", "value string")

	return headers

}
