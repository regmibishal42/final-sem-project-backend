package auth_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
)

func (r AuthRepository) CreateUser(ctx context.Context, input model.UserInput) *model.AuthMutationResponse {
	user := model.User{
		UserType: model.UserTypeAdmin,
	}

	//validate email
	err := util.IsEmailValid(input.Email)
	if err != nil {
		return &model.AuthMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.VALIDATION_ERROR, util.Ref("email")),
		}
	}
	user.Email = input.Email
	//hash password
	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return &model.AuthMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}
	user.Password = hashedPassword

	err = r.TableUser.CreateUser(ctx, &user)
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

func (r AuthRepository) GetUserByID(ctx context.Context, userID *string) (*model.User, error) {
	user, err := r.TableUser.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil

}
