package auth_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
)

func (r AuthRepository) CreateUser(ctx context.Context, input model.UserInput) *model.AuthMutationResponse {
	user, validationError := input.Validator()
	if validationError != nil {
		return &model.AuthMutationResponse{
			Data:  nil,
			Error: validationError,
		}
	}
	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return &model.AuthMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}
	user.Password = hashedPassword

	//create user
	user.Profile = &model.Profile{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	err = r.TableUser.CreateUser(ctx, user)
	if err != nil {
		return &model.AuthMutationResponse{
			Data:  nil,
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}
	//send otp to user
	_, err = r.CreateOtp(ctx, user)
	if err != nil {
		return &model.AuthMutationResponse{
			Data:  nil,
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}

	return &model.AuthMutationResponse{
		Data:  user,
		Error: nil,
	}
}

func (r AuthRepository) GetUserByID(ctx context.Context, userID *string) (*model.User, error) {
	return r.TableUser.GetUserByID(ctx, userID)
}

func (r AuthRepository) GetUserDetailsByID(ctx context.Context, userID *string) *model.AuthQueryResponse {
	user, err := r.TableUser.GetUserByID(ctx, userID)
	if err != nil {
		return &model.AuthQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}
	return &model.AuthQueryResponse{
		Data: []*model.User{user},
	}
}

//verify user
func (r AuthRepository) VerifyUser(ctx context.Context, input *model.UserVerificationInput) *model.AuthMutationResponse {
	user, err := r.TableUser.GetUserByID(ctx, &input.UserID)
	if err != nil {
		return &model.AuthMutationResponse{
			Data:  nil,
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}
	if user.IsVerified {
		return &model.AuthMutationResponse{
			Data:  nil,
			Error: exception.MutationErrorHandler(ctx, errors.New("user already verified"), exception.BAD_REQUEST, nil),
		}
	}
	//get otp
	otp, err := r.TableOtp.GetOtp(ctx, user.ID)
	if err != nil {
		return &model.AuthMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}
	//todo: check if otp is expired

	//validate opt
	if otp != nil && otp.Secret == input.Otp {
		//update user verification
		user.IsVerified = true
		err := r.TableUser.UpdateUserDetails(ctx, user)
		if err != nil {
			return &model.AuthMutationResponse{
				Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
			}
		}
		return &model.AuthMutationResponse{
			Data: user,
		}

	}
	return &model.AuthMutationResponse{
		Data:  nil,
		Error: exception.MutationErrorHandler(ctx, errors.New("invalid otp"), exception.BAD_REQUEST, nil),
	}
}
