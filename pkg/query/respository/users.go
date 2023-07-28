package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateUser(ctx context.Context, user *model.User) error {
	err := r.db.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
