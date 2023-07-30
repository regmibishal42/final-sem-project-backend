package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"backend/exception"
	"backend/graph/generated"
	"backend/graph/model"
	"context"
)

// Profile is the resolver for the profile field.
func (r *userResolver) Profile(ctx context.Context, obj *model.User) (*model.Profile, error) {
	return r.AuthDomain.GetProfileByUserID(ctx, obj.ID)
}

// CreateUser is the resolver for the createUser field.
func (r *userMutationResolver) CreateUser(ctx context.Context, obj *model.UserMutation, input model.UserInput) (*model.AuthMutationResponse, error) {
	return r.AuthDomain.CreateUser(ctx, input), nil
}

// LoginUser is the resolver for the loginUser field.
func (r *userMutationResolver) LoginUser(ctx context.Context, obj *model.UserMutation, input model.LoginInput) (*model.AuthResponse, error) {
	return r.AuthDomain.Login(ctx, &input), nil
}

// VerifyUser is the resolver for the verifyUser field.
func (r *userMutationResolver) VerifyUser(ctx context.Context, obj *model.UserMutation, input model.UserVerificationInput) (*model.AuthMutationResponse, error) {
	return r.AuthDomain.VerifyUser(ctx, &input), nil
}

// GetUserDetails is the resolver for the getUserDetails field.
func (r *userQueryResolver) GetUserDetails(ctx context.Context, obj *model.UserQuery, input *model.GetUserInput) (*model.AuthQueryResponse, error) {
	if input != nil {
		return r.AuthDomain.GetUserDetailsByID(ctx, &input.ID), nil
	}
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.AuthQueryResponse{
			Data:  nil,
			Error: exception.QueryErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	return r.AuthDomain.GetUserDetailsByID(ctx, &user.ID), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserMutation returns generated.UserMutationResolver implementation.
func (r *Resolver) UserMutation() generated.UserMutationResolver { return &userMutationResolver{r} }

// UserQuery returns generated.UserQueryResolver implementation.
func (r *Resolver) UserQuery() generated.UserQueryResolver { return &userQueryResolver{r} }

type userResolver struct{ *Resolver }
type userMutationResolver struct{ *Resolver }
type userQueryResolver struct{ *Resolver }
