package query_controller

import (
	"backend/graph/model"
	"context"
)

type StaffQueryInterface interface {
	CreateStaff(ctx context.Context, staff *model.Staff) error
}
