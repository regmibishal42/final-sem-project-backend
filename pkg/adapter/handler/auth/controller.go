package auth_handler

import (
	"backend/graph/model"
	"context"
)

type AuthController struct {
	AuthDomain AuthInterface
}

type AuthInterface interface {
	CreateUser(ctx context.Context, input model.UserInput) *model.AuthMutationResponse
	GetUserByID(ctx context.Context, userID *string) (*model.User, error)

	//login
	Login(ctx context.Context, input *model.LoginInput) *model.AuthResponse
}
