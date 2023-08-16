package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateSales(ctx context.Context, sales *model.Sales) error {
	err := r.db.Model(&model.Sales{}).Create(&sales).Error
	if err != nil {
		return err
	}
	return nil
}
