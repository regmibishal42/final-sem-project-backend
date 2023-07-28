package auth_handler

import (
	"backend/exception"
	"backend/graph/model"
	"context"
)

func (r AuthRepository) CreateUser(ctx context.Context, input model.UserInput) *model.AuthMutationResponse {
	user := model.User{
		Email:    input.Email,
		UserType: model.UserTypeAdmin,
	}
	err := r.TableUser.CreateUser(ctx, &user)
	if err != nil {
		return &model.AuthMutationResponse{
			Data:  nil,
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}

	return &model.AuthMutationResponse{
		Data:  &user,
		Error: nil,
	}
}
