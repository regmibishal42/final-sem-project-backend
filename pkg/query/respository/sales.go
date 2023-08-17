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

func (r QueryRepository) GetSalesByFilter(ctx context.Context, filter *model.FilterSalesInput, organizationID *string, pageInfo *model.OffsetPageInfo) ([]*model.Sales, error) {
	sales := []*model.Sales{}
	now := time.Now()
	db := r.db.Model(&model.Sales{}).Where("sales.deleted_at IS NULL AND sales.organization_id = ?", organizationID)
	if filter.Params != nil {
		if filter.Params.FilterType == model.SalesInfoTypeDaily {
			today := time.Now().Truncate(24 * time.Hour)
			// Truncate to beginning of the day
			db = db.Where("sales.created_at >= ?", today)
		}
		if filter.Params.FilterType == model.SalesInfoTypeWeekly {
			beginningOfWeek := now.AddDate(0, 0, int(time.Sunday-now.Weekday()))
			db = db.Where("sales.created_at >= ?", beginningOfWeek)
		}
		if filter.Params.FilterType == model.SalesInfoTypeMonthly {
			beginningOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			db = db.Where("sales.created_at >= ?", beginningOfMonth)
		}
		if filter.Params.FilterType == model.SalesInfoTypeYearly {
			beginningOfYear := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
			db = db.Where("sales.created_at >= ?", beginningOfYear)
		}
		if filter.Params.ProductID != nil {
			db = db.Where("sales.product_id = ?", filter.Params.ProductID)
		}
		if filter.Params.CategoryID != nil {
			db = db.Joins("left join products on sales.product_id = products.id").Where("products.category_id = ?", filter.Params.CategoryID)
		}
	}
	err := db.Scopes(Paginate(sales, pageInfo, filter.Page, db)).Find(&sales).Error
	//err := db.Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (r QueryRepository) GetSalesByID(ctx context.Context, salesID *string, organizationID *string) (*model.Sales, error) {
	sales := model.Sales{}
	err := r.db.Where("deleted_at IS NULL AND id = ? AND organization_id = ?", salesID, organizationID).Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return &sales, nil
}
