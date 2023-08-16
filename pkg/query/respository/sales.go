package query_repository

import (
	"backend/graph/model"
	"context"
	"time"

	"gorm.io/gorm/clause"
)

func (r QueryRepository) CreateSales(ctx context.Context, sales *model.Sales) error {
	err := r.db.Model(&model.Sales{}).Create(&sales).Error
	if err != nil {
		return err
	}
	return nil
}

func (r QueryRepository) UpdateSales(ctx context.Context, sales *model.Sales, organizationID *string) (*model.Sales, error) {
	updatedSales := model.Sales{}
	err := r.db.Model(&model.Sales{}).Clauses(clause.Returning{}).Where("deleted_at IS NULL AND id = ? AND organization_id = ?", sales.ID, organizationID).Updates(&sales).Find(&updatedSales).Error
	if err != nil {
		return nil, err
	}
	return &updatedSales, nil
}

func (r QueryRepository) DeleteSales(ctx context.Context, salesID *string, organizationID *string) error {
	err := r.db.Model(&model.Sales{}).Where("deleted_at IS NULL AND id = ? AND organization_id = ?", salesID, organizationID).Update("deleted_at", time.Now()).Error
	return err
}
