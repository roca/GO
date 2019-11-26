package utils

import (
	"net/http"
)

// GetUserID return the User ID in the request header
func GetUserID(header http.Header) string {
	return header.Get("app_user_id")
}

// GetUserName return the User Name in the request header
func GetUserName(header http.Header) string {
	return header.Get("app_user_name")
}

// GetResponseHeaders return default response headers
func GetResponseHeaders() map[string]string {

	headers := make(map[string]string)

	headers["key string"] = "value string"

	return headers

}
