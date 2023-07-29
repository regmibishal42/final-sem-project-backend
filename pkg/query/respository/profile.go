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
