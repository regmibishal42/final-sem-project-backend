package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"backend/graph/generated"
	"backend/graph/model"
	"context"
	"fmt"
)

// Address is the resolver for the Address field.
func (r *profileResolver) Address(ctx context.Context, obj *model.Profile) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: Address - Address"))
}

// CreateProfile is the resolver for the createProfile field.
func (r *profileMutationResolver) CreateProfile(ctx context.Context, obj *model.ProfileMutation, input model.CreateProfileInput) (*model.ProfileMutationResponse, error) {
	panic(fmt.Errorf("not implemented: CreateProfile - createProfile"))
}

// UpdateProfile is the resolver for the updateProfile field.
func (r *profileMutationResolver) UpdateProfile(ctx context.Context, obj *model.ProfileMutation, input model.UpdateProfileInput) (*model.ProfileMutationResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateProfile - updateProfile"))
}

// GetProfile is the resolver for the getProfile field.
func (r *profileQueryResolver) GetProfile(ctx context.Context, obj *model.ProfileQuery, input *model.GetByIDInput) (*model.ProfileQueryResponse, error) {
	panic(fmt.Errorf("not implemented: GetProfile - getProfile"))
}

// Profile returns generated.ProfileResolver implementation.
func (r *Resolver) Profile() generated.ProfileResolver { return &profileResolver{r} }

// ProfileMutation returns generated.ProfileMutationResolver implementation.
func (r *Resolver) ProfileMutation() generated.ProfileMutationResolver {
	return &profileMutationResolver{r}
}

// ProfileQuery returns generated.ProfileQueryResolver implementation.
func (r *Resolver) ProfileQuery() generated.ProfileQueryResolver { return &profileQueryResolver{r} }

type profileResolver struct{ *Resolver }
type profileMutationResolver struct{ *Resolver }
type profileQueryResolver struct{ *Resolver }
