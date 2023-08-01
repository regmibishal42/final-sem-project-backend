package query_controller

import (
	"backend/graph/model"
	"context"
)

type OrganizationQueryInterface interface {
	CreateOrganization(ctx context.Context, organization *model.Organization) error
}
