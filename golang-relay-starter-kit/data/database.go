package data

// Post Data model structs
type Post struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author Author
}

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Mock data
var latestPost = &Post{
	ID:     "1",
	Text:   "Hello World",
	Author: Author{ID: "2", Name: "Ray Bradbury"},
}

// GetLatestPost Data getters/setters
func GetLatestPost() *Post {
	return latestPost
}
