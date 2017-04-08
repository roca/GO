package data

// Post Data model structs
type Post struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

// Mock data
var latestPost = &Post{
	ID:     "1",
	Text:   "Hello World",
	Author: "Ray Bradbury",
}

// GetLatestPost Data getters/setters
func GetLatestPost() *Post {
	return latestPost
}
