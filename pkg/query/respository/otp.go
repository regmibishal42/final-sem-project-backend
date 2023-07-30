package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateOtp(ctx context.Context, otp *model.Otp) error {
	err := r.db.Model(&model.Otp{}).Create(&otp).Error
	if err != nil {
		return err
	}
	return nil
}

func (r QueryRepository) GetOtp(ctx context.Context, userID string) (*model.Otp, error) {
	otp := model.Otp{}
	err := r.db.Where("user_id = ?", userID).Find(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}
