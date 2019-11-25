package utils

// GetResponseHeaders return default response headers
func GetResponseHeaders() map[string]string {

	header := make(map[string]string)

	header["key string"] = "value string"

	return header

}
