package handlers

// API struct
type API struct {
	Message string "json:message"
}

// User struct
type User struct {
	ID    int    `json:id`
	Name  string `json:username`
	Email string `json:email`
	First string `json:first`
	Last  string `json:last`
}
