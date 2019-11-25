package utils

// GetResponseHeaders return default response headers
func GetResponseHeaders() map[string]string {

	headers := make(map[string]string)

	headers["key string"] = "value string"

	return headers

}
