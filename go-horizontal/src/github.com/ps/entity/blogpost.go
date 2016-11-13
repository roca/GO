package entity

import "time"

type ContentItem struct {
	ID          int
	Subject     string
	Body        string
	Author      *Author
	Comments    []Comment
	CreatedDate *time.Time
	PublishDate *time.Time
	IsPublished bool
}

type BlogPost struct {
	ContentItem
}
