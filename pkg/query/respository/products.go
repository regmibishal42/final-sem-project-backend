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

func (r QueryRepository) UpdateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	updatedProduct := model.Product{}
	err := r.db.Model(&model.Product{}).Where("deleted_at IS NULL AND id = ?", product.ID).Updates(&product).Find(&updatedProduct).Error
	if err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}
