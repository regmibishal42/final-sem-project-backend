package auth_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
	"time"
)

func (r AuthRepository) CreateOtp(ctx context.Context, user *model.User) (*model.Otp, error) {
	otp := model.Otp{
		UserId:    user.ID,
		CreatedAt: time.Now(),
		Secret:    util.OtpGenerator(),
	}
	err := r.TableOtp.CreateOtp(ctx, &otp)
	if err != nil {
		return nil, err
	}
	//send otp email to user
	go util.SendOtpEmail("User", user.Email, otp.Secret)
	return &otp, nil

}

//update otp for resend functionality
func (r AuthRepository) UpdateOtp(ctx context.Context, userID *string) *model.OtpMutationResponse {
	//validate the id from input
	if !util.IsValidID(*userID) {
		return &model.OtpMutationResponse{
			Data:  util.Ref(false),
			Error: exception.MutationErrorHandler(ctx, errors.New("invalid ID"), exception.BAD_REQUEST, nil),
		}
	}
	//get user from id
	user, err := r.TableUser.GetUserByID(ctx, userID)
	if err != nil {
		return &model.OtpMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.NOT_FOUND, nil),
		}
	}
	//get otp
	otpSecret := util.OtpGenerator()
	//update the otp
	otp := model.Otp{
		UserId:    *userID,
		UpdatedAt: util.Ref(time.Now()),
		Secret:    otpSecret,
	}
	err = r.TableOtp.UpdateOtp(ctx, &otp)
	if err != nil {
		return &model.OtpMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}
	}
	go util.SendOtpEmail("User", user.Email, otp.Secret)
	return &model.OtpMutationResponse{
		Data:  util.Ref(true),
		Error: nil,
	}

}

//verify the otp
func (r AuthRepository) VerifyOtp(ctx context.Context, input *model.VerifyOtpInput) *model.OtpMutationResponse {
	user, err := r.TableUser.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return &model.OtpMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.NOT_FOUND, nil),
		}
	}
	if user == nil {
		return &model.OtpMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("user with given email not found"), exception.NOT_FOUND, nil),
		}
	}
	//get otp
	otp, err := r.TableOtp.GetOtp(ctx, user.ID)
	if err != nil {
		return &model.OtpMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.NOT_FOUND, nil),
		}
	}
	if input.Otp != otp.Secret {
		return &model.OtpMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("opt not matched"), exception.BAD_REQUEST, nil),
		}
	}
	return &model.OtpMutationResponse{
		Data:  util.Ref(true),
		Error: nil,
	}

}
