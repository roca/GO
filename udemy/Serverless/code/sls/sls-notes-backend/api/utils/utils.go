package utils

// GetUserID return the User ID in the request header
func GetUserID(headers map[string]string) string {
	return headers["app_user_id"]
}

// GetUserName return the User Name in the request header
func GetUserName(headers map[string]string) string {
	return headers["app_user_name"]
}

// GetResponseHeaders return default response headers
func GetResponseHeaders() map[string]string {

	headers := make(map[string]string)

	headers["Access-Control-Allow-Origin"] = "*"

	return headers

}
