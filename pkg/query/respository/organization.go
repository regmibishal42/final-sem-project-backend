package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateOrganization(ctx context.Context, organization *model.Organization) error {
	err := r.db.Model(&model.Organization{}).Create(&organization).Error
	if err != nil {
		return err
	}
	return nil
}
