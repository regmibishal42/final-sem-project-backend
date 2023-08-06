package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateStaff(ctx context.Context, staff *model.Staff) error {
	err := r.db.Model(model.Staff{}).Create(&staff).Error
	if err != nil {
		return err
	}
	return nil
}
