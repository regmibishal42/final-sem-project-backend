package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateProduct(ctx context.Context, product *model.Product) error {
	err := r.db.Model(&model.Product{}).Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}
