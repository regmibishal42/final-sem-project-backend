package query_controller

import (
	"backend/graph/model"
	"context"
)

type OtpQueryInterface interface {
	CreateOtp(ctx context.Context, otp *model.Otp) error
	// UpdateOtp(ctx context.Context, otp *model.Otp) error
	GetOtp(ctx context.Context, userID string) (*model.Otp, error)
	UpdateOtp(ctx context.Context, otp *model.Otp) error
}
