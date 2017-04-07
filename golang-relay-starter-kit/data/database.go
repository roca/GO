package data

// Post Data model structs
type Post struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// Mock data
var latestPost = &Post{
	ID:   "1",
	Text: "Hello World",
}

// GetLatestPost Data getters/setters
func GetLatestPost() *Post {
	return latestPost
}
