package query_controller

import (
	"backend/graph/model"
	"context"
)

type StaffQueryInterface interface {
	CreateStaff(ctx context.Context, staff *model.Staff) error
	UpdateStaff(ctx context.Context, staff *model.Staff) error
	GetStaffByID(ctx context.Context, staffID *string) (*model.Staff, error)
	GetStaffsByOrganization(ctx context.Context, organizationID *string) ([]*model.Staff, error)
}
