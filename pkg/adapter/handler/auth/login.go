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
