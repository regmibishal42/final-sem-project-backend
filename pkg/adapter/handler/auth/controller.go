package auth_handler

import (
	"backend/graph/model"
	"context"
)

type AuthController struct {
	AuthDomain AuthInterface
}

type AuthInterface interface {
	CreateUser(ctx context.Context, input model.UserInput) *model.AuthMutationResponse
	GetUserByID(ctx context.Context, userID *string) (*model.User, error)
	GetUserDetailsByID(ctx context.Context, userID *string) *model.AuthQueryResponse
	VerifyUser(ctx context.Context, input *model.UserVerificationInput) *model.AuthMutationResponse
	UpdateUserPassword(ctx context.Context, user *model.User, input *model.UpdatePasswordInput) (*model.AuthMutationResponse, error)
	ForgetUserPassword(ctx context.Context, input *model.ForgetPasswordInput) (*model.RegisterResponse, error)
	ResetPassword(ctx context.Context, input *model.ResetPasswordInput) (*model.RegisterResponse, error)

	//login
	Login(ctx context.Context, input *model.LoginInput) *model.AuthResponse

	//Profile
	GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error)

	//otp
	CreateOtp(ctx context.Context, user *model.User) (*model.Otp, error)
	UpdateOtp(ctx context.Context, userID *string) *model.OtpMutationResponse
	VerifyOtp(ctx context.Context, otp *model.VerifyOtpInput) *model.OtpMutationResponse
}
