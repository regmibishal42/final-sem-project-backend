package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateUserProfile(ctx context.Context, profile *model.Profile) error {
	err := r.db.Model(&model.Profile{}).Create(&profile).Error
	if err != nil {
		return err
	}
	return nil
}

func (r QueryRepository) GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error) {
	userProfile := model.Profile{}
	err := r.db.Model(&model.Profile{}).Where("user_id = ?", userID).Find(&userProfile).Error
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}
