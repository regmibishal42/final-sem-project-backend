package auth_handler

import (
	"backend/graph/model"
	"context"
)

func (r AuthRepository) CreateUser(ctx context.Context, input model.UserInput) *model.AuthMutationResponse {
	return &model.AuthMutationResponse{
		Data: &model.User{
			Email: "Hello@Success",
		},
	}
}
