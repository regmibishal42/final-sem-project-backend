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

func (r QueryRepository) DeleteProduct(ctx context.Context, productID *string) error {
	product := model.Product{}
	deletedProduct := model.DeletedProducts{}
	tx := r.db.Begin()
	if err := tx.Where("deleted_at IS NULL AND id = ?", productID).Find(&product).Error; err != nil {
		tx.Rollback()
		return err
	}
	deletedProduct.Product = &product
	if err := tx.Delete(&product).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&deletedProduct).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r QueryRepository) GetProductByID(ctx context.Context, productID *string) (*model.Product, error) {
	product := model.Product{}
	err := r.db.Model(&model.Product{}).Where("deleted_at IS NULL AND id = ?", productID).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
