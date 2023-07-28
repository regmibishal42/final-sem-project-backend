package query_controller

import (
	"backend/graph/model"
	"context"
)

type UserQueryInterface interface {
	CreateUser(ctx context.Context, user *model.User) error
}
