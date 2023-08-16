package query_repository

import (
	"backend/graph/model"
	"context"

	"gorm.io/gorm/clause"
)

func (r QueryRepository) CreateSales(ctx context.Context, sales *model.Sales) error {
	err := r.db.Model(&model.Sales{}).Create(&sales).Error
	if err != nil {
		return err
	}
	return nil
}

func (r QueryRepository) UpdateSales(ctx context.Context, sales *model.Sales) (*model.Sales, error) {
	updatedSales := model.Sales{}
	err := r.db.Model(&model.Sales{}).Clauses(clause.Returning{}).Where("deleted_at IS NULL AND id = ?", sales.ID).Updates(&sales).Find(&updatedSales).Error
	if err != nil {
		return nil, err
	}
	return &updatedSales, nil
}
