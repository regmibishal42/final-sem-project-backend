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
	GetOrganizationByID(ctx context.Context, input *model.OrganizationInput) (*model.OrganizationQueryResponse, error)
	GetOrganizationDetailsByID(ctx context.Context, organizationID string) (*model.Organization, error)

	GetOrganizationByFilter(ctx context.Context, filters *model.OrganizationFilterInput) (*model.OrganizationsQueryResponse, error)

	// Staff
	CreateStaff(ctx context.Context, user *model.User, input *model.CreateStaffInput) (*model.StaffMutationResponse, error)
	UpdateStaffDetails(ctx context.Context, user *model.User, input *model.UpdateStaffInput) (*model.StaffMutationResponse, error)
	GetStaffByID(ctx context.Context, user *model.User, input *model.GetStaffInput) (*model.StaffQueryResponse, error)
	GetStaffsByOrganization(ctx context.Context, user *model.User, input *model.GetOrganizationStaffsInput) (*model.StaffsQueryResponse, error)

	//shared
	GetOrganizationIDFromUserID(ctx context.Context, user *model.User) (*string, error)
}
