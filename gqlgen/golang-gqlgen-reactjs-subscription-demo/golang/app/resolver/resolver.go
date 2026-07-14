package resolver

import (
	"context"

	"github.com/GOCODE/gqlgen/golang-gqlgen-reactjs-subscription-demo/golang/app/model"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddChannel(ctx context.Context, name string) (model.Channel, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateChannel(ctx context.Context, id int, name string) (model.Channel, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteChannel(ctx context.Context, ID int) (model.Channel, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Channels(ctx context.Context) ([]model.Channel, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) SubscriptionChannelAdded(ctx context.Context) (<-chan model.Channel, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) SubscriptionChannelDeleted(ctx context.Context) (<-chan model.Channel, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) SubscriptionChannelUpdated(ctx context.Context) (<-chan model.Channel, error) {
	panic("not implemented")
}
