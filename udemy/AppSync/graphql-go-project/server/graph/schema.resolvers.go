package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/roca/gqlgen-posts/graph/generated"
	"github.com/roca/gqlgen-posts/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post := &model.Post{
		ID:      fmt.Sprintf("T%d", rand.Int()),
		Comment: input.Comment,
		UserID:  input.UserID,
	}
	r.posts = append(r.posts, post)

	for _, user := range r.users {
		if user.ID == input.UserID {
			user.Posts = append(user.Posts, post)
		}
	}

	return post, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:         fmt.Sprintf("T%d", rand.Int()),
		Name:       input.Name,
		Age:        input.Age,
		Profession: input.Profession,
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == obj.UserID {
			return user, nil
		}
	}

	return nil, errors.New("User not found")
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.posts, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("Post not found")
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	for _, post := range r.posts {
		if post.ID == id {
			return post, nil
		}
	}
	return nil, errors.New("Post not found")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
