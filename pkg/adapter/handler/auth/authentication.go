package auth_handler

import (
	"backend/graph/model"
	"context"
	"time"
)

func (r AuthRepository) CreateUser(ctx context.Context, input model.UserInput) *model.AuthMutationResponse {
	user := model.User{
		Email:     input.Email,
		UserType:  model.UserTypeAdmin,
		CreatedAt: time.Now(),
	}
	err := r.TableUser.CreateUser(ctx, &user)
	if err != nil {
		return &model.AuthMutationResponse{
			Data:  nil,
			Error: model.ServerError{Message: err.Error(), Code: 500},
		}
	}

	return &model.AuthMutationResponse{
		Data:  &user,
		Error: nil,
	}
}
