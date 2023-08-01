package query_controller

import (
	"backend/graph/model"
	"context"
)

type OrganizationQueryInterface interface {
	CreateOrganization(ctx context.Context, organization *model.Organization) error
	GetOrganizationByID(ctx context.Context, organizationID *string) (*model.Organization, error)
	GetOrganizationsByFilter(ctx context.Context, filter *model.OrganizationFilterInput) ([]*model.Organization, error)
}
