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

func (r QueryRepository) GetStaffByID(ctx context.Context, staffID *string) (*model.Staff, error) {
	staff := model.Staff{}
	err := r.db.Model(&model.Staff{}).Where("staff_id = ?", staffID).First(&staff).Error
	if err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r QueryRepository) GetStaffsByOrganization(ctx context.Context, organizationID *string) ([]*model.Staff, error) {
	staffs := []*model.Staff{}
	err := r.db.Model(&model.Staff{}).Where("organization_id = ?", organizationID).Find(&staffs).Error
	if err != nil {
		return nil, err
	}
	return staffs, nil
}
