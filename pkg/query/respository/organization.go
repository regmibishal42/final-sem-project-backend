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

func (r QueryRepository) GetOrganizationByID(ctx context.Context, organizationID *string) (*model.Organization, error) {
	organization := model.Organization{}
	err := r.db.Where("deleted_at IS NULL AND id = ?", organizationID).First(&organization).Error
	if err != nil {
		return nil, err
	}
	return &organization, nil
}

func (r QueryRepository) GetOrganizationsByFilter(ctx context.Context, filter *model.OrganizationFilterInput) ([]*model.Organization, error) {
	organizations := []*model.Organization{}
	db := r.db.Model(&model.Organization{}).Where("deleted_at IS NULL")
	if filter != nil {
		if filter.VerificationStatus != nil {
			db = db.Where("verification_status = ?", filter.VerificationStatus)
		}
	}
	err := db.Find(&organizations).Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r QueryRepository) GetOrganizationIDByUser(ctx context.Context, user *model.User) (*string, error) {
	var id string
	db := r.db.Model(&model.Organization{})
	if user.UserType == model.UserTypeAdmin {
		err := db.Where("created_by_id = ?", user.ID).Select("id").Find(&id).Error
		if err != nil {
			return nil, err
		}
		return &id, nil
	}
	err := db.Joins("left join staffs on staffs.organization_id = organizations.id").Where("staffs.staff_id = ?", user.ID).Select("organizations.id").Find(&id).Error
	if err != nil {
		return nil, err
	}
	return &id, nil
}
