package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateCategory(ctx context.Context, category *model.Category) error {
	err := r.db.Model(&model.Category{}).Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}
