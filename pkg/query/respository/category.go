package query_repository

import (
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
	"time"

	"gorm.io/gorm/clause"
)

func (r QueryRepository) CreateCategory(ctx context.Context, category *model.Category) error {
	err := r.db.Model(&model.Category{}).Create(&category).Error
	if err != nil {

		if util.IsDuplicateKeyError(err) {
			dbRes := r.db.Model(&model.Category{}).Where("name = ? AND organization_id = ? AND deleted_at IS NOT NULL", category.Name, category.OrganizationID).
				Update("deleted_at", nil)
			if dbRes.Error != nil {
				return err
			}
			if dbRes.RowsAffected < 1 {
				return errors.New("category already exist")
			}
			return nil
		}
		return err
	}
	return nil
}

func (r QueryRepository) GetCategoryByOrganization(ctx context.Context, organizationID *string) ([]*model.Category, error) {
	categories := []*model.Category{}
	err := r.db.Model(&model.Category{}).Where("deleted_at IS NULL AND organization_id = ?", organizationID).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r QueryRepository) DeleteCategory(ctx context.Context, categoryID *string) error {
	err := r.db.Model(&model.Category{}).Where("id = ?", categoryID).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}

func (r QueryRepository) GetCategoryByID(ctx context.Context, categoryID *string) (*model.Category, error) {
	category := model.Category{}
	err := r.db.Model(&model.Category{}).Clauses(clause.Returning{}).Where("deleted_at IS NULL AND id = ?", categoryID).Find(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
