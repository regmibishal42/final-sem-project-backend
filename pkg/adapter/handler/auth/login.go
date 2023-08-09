package auth_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"
)

func (r AuthRepository) Login(ctx context.Context, input *model.LoginInput) *model.AuthResponse {
	//input validation
	if validationError := input.Validator(); validationError != nil {
		return &model.AuthResponse{
			Error: validationError,
		}
	}
	email := strings.ToLower(input.Email)
	user, err := r.TableUser.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.AuthResponse{
				Error: exception.MutationErrorHandler(ctx, errors.New("invalid email or password"), exception.BAD_REQUEST, nil),
			}
		}
		return &model.AuthResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.BAD_REQUEST, nil),
		}
	}
	isPasswordMatch := util.CheckPasswordHash(input.Password, user.Password)
	if !isPasswordMatch {
		return &model.AuthResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("invalid email or password"), exception.BAD_REQUEST, nil),
		}
	}
	token, err := user.Token()
	if err != nil {
		return &model.AuthResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}
	return &model.AuthResponse{
		Data: token,
	}
}

func (r AuthRepository) UpdateUserPassword(ctx context.Context, user *model.User, input *model.UpdatePasswordInput) (*model.AuthMutationResponse, error) {
	if validationError := input.Validator(); validationError != nil {
		return &model.AuthMutationResponse{
			Error: validationError,
		}, nil
	}
	//compare old password
	isValidOldPassword := util.CheckPasswordHash(input.OldPassword, user.Password)
	if !isValidOldPassword {
		return &model.AuthMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("invalid old password"), exception.BAD_REQUEST, nil),
		}, nil
	}
	//update password
	hashedNewPassword, err := util.HashPassword(input.NewPassword)
	if err != nil {
		return &model.AuthMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	user.Password = hashedNewPassword
	err = r.TableUser.UpdateUserDetails(ctx, user)
	if err != nil {
		return &model.AuthMutationResponse{}, nil
	}
	return &model.AuthMutationResponse{
		Data: user,
	}, nil
}

func (r AuthRepository) ForgetUserPassword(ctx context.Context, input *model.ForgetPasswordInput) (*model.RegisterResponse, error) {
	//validate input
	if validationError := input.ValidationError(); validationError != nil {
		return &model.RegisterResponse{
			Error: validationError,
		}, nil
	}
	//get user from email
	user, err := r.TableUser.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return &model.RegisterResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.RegisterResponse{
		UserID: &user.ID,
	}, nil

}

//reset password -> after user email is verified via forget password
func (r AuthRepository) ResetPassword(ctx context.Context, input *model.ResetPasswordInput) (*model.RegisterResponse, error) {
	if validationError := input.Validator(); validationError != nil {
		return &model.RegisterResponse{
			Error: validationError,
		}, nil
	}
	//get user from email
	user, err := r.TableUser.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return &model.RegisterResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	//update the password
	hashedPassword, err := util.HashPassword(input.NewPassword)
	if err != nil {
		return &model.RegisterResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	user.Password = hashedPassword
	err = r.TableUser.UpdateUserDetails(ctx, user)
	if err != nil {
		return &model.RegisterResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.RegisterResponse{
		UserID: &user.ID,
	}, nil
}
