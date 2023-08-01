package organization_handler

import (
	"backend/graph/model"
	"context"
)

type OrganizationController struct {
	OrganizationDomain OrganizationInterface
}

type OrganizationInterface interface {
	CreateOrganization(ctx context.Context, input *model.CreateOrganizationInput, user *model.User) (*model.OrganizationMutationResponse, error)
}
