package query_repository

import (
	"backend/graph/model"
	"context"
	"strings"
)

func (r QueryRepository) CreateProduct(ctx context.Context, product *model.Product) error {
	err := r.db.Model(&model.Product{}).Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r QueryRepository) UpdateProduct(ctx context.Context, product *model.Product, organizationID *string) (*model.Product, error) {
	updatedProduct := model.Product{}
	err := r.db.Model(&model.Product{}).Where("deleted_at IS NULL AND id = ? AND organization_id = ?", product.ID, organizationID).Updates(&product).Find(&updatedProduct).Error
	if err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}

func (r QueryRepository) DeleteProduct(ctx context.Context, productID *string, organizationID *string) error {
	product := model.Product{}
	deletedProduct := model.DeletedProducts{}
	tx := r.db.Begin()
	if err := tx.Where("deleted_at IS NULL AND id = ? AND organization_id = ?", productID, organizationID).Find(&product).Error; err != nil {
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

func (r QueryRepository) GetProductsByFilter(ctx context.Context, pageInfo *model.OffsetPageInfo, filter *model.GetProductsByFilterInput, organizationID *string) ([]*model.Product, error) {
	products := []*model.Product{}
	db := r.db.Model(&model.Product{}).Where("deleted_at IS NULL AND organization_ID = ?", organizationID)
	if filter.Params != nil {
		if filter.Params.CategoryID != nil {
			db = db.Where("category_id = ?", filter.Params.CategoryID)
		}
		if filter.Params.SearchQuery != nil && *filter.Params.SearchQuery != "" {
			splitQuery := strings.Split(strings.TrimSpace(*filter.Params.SearchQuery), " ")
			db = db.Where("LOWER(name) LIKE(?) OR lower(name) LIKE(?)", "%"+*filter.Params.SearchQuery+"%", "%"+splitQuery[0]+"%")
		}
	}
	err := db.Scopes(Paginate(products, pageInfo, filter.Page, db)).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
