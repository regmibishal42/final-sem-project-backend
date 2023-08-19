package query_controller

import (
	"backend/graph/model"
	"context"
)

type UserQueryInterface interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, userID *string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUserDetails(ctx context.Context, user *model.User) error
	GetAdditionalInformation(ctx context.Context, userID *string) (*model.AdditionalUserInformation, error)
}
