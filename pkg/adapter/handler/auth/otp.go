package auth_handler

import (
	"backend/graph/model"
	"backend/pkg/util"
	"context"
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
	go util.SendVerificationEmail(otp.Secret, user.Email)
	return &otp, nil

}
