package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"backend/graph/generated"
	"backend/graph/model"
	"context"
)

// Auth is the resolver for the auth field.
func (r *mutationResolver) Auth(ctx context.Context) (*model.UserMutation, error) {
	return &model.UserMutation{}, nil
}

// Profile is the resolver for the profile field.
func (r *mutationResolver) Profile(ctx context.Context) (*model.ProfileMutation, error) {
	return &model.ProfileMutation{}, nil
}

// Auth is the resolver for the auth field.
func (r *queryResolver) Auth(ctx context.Context) (*model.UserQuery, error) {
	return &model.UserQuery{}, nil
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context) (*model.ProfileQuery, error) {
	return &model.ProfileQuery{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
