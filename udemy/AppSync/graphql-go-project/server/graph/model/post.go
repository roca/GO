package model

type Post struct {
	ID      string `json:"id"`
	Comment string `json:"comment"`
	UserID  string `json:"user"`
}
