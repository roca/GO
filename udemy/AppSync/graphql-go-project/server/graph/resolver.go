package graph

import "github.com/roca/gqlgen-posts/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	posts []*model.Post
	users []*model.User
	user  *model.User
	post  *model.Post
}
