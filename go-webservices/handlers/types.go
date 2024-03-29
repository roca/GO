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

type Users struct {
	Users []User `json:"users"`
}

type CreateResponse struct {
	Error     string `json:error`
	ErrorCode int    `json:code`
}

type ErrMsg struct {
	ErrCode    int
	StatusCode int
	Msg        string
}

type Page struct {
	Title string
	Body  []byte
}
