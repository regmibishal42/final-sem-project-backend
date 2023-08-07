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

func (r QueryRepository) UpdateStaff(ctx context.Context, staff *model.Staff) error {
	err := r.db.Model(&model.Staff{}).Where("staff_id = ?", staff.StaffID).Updates(&staff).Error
	if err != nil {
		return err
	}
	return nil
}
