package organization_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
	"fmt"
)

func (r OrganizationRepository) CreateOrganization(ctx context.Context, input *model.CreateOrganizationInput, user *model.User) (*model.OrganizationMutationResponse, error) {
	organization, validationError := input.Validator()
	if validationError != nil {
		return &model.OrganizationMutationResponse{
			Error: validationError,
		}, nil
	}
	//check if user has permission/privilege
	if user.UserType != model.UserTypeLogicloud {
		if user.UserType != model.UserTypeAdmin {
			return &model.OrganizationMutationResponse{
				Error: exception.MutationErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
	}

	//	organization.CreatedBy = user
	organization.CreatedByID = user.ID
	fmt.Println("User ID is", user.ID)
	//create organization
	err := r.TableOrganization.CreateOrganization(ctx, &organization)
	if err != nil {
		return &model.OrganizationMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}

	return &model.OrganizationMutationResponse{
		Data: &organization,
	}, nil
}
func (r OrganizationRepository) GetOrganizationByID(ctx context.Context, input *model.OrganizationInput) (*model.OrganizationQueryResponse, error) {
	//validate id from the input
	if !util.IsValidID(input.ID) {
		return &model.OrganizationQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("invalid organization id"), exception.BAD_REQUEST, nil),
		}, nil
	}
	organization, err := r.TableOrganization.GetOrganizationByID(ctx, &input.ID)
	if err != nil {
		return &model.OrganizationQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.OrganizationQueryResponse{
		Data: organization,
	}, nil
}

func (r OrganizationRepository) GetOrganizationByFilter(ctx context.Context, filters *model.OrganizationFilterInput) (*model.OrganizationsQueryResponse, error) {
	organizations, err := r.TableOrganization.GetOrganizationsByFilter(ctx, filters)
	if err != nil {
		return &model.OrganizationsQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.OrganizationsQueryResponse{
		Data: organizations,
	}, nil
}

//get organization only
func (r OrganizationRepository) GetOrganizationDetailsByID(ctx context.Context, organizationID string) (*model.Organization, error) {
	return r.TableOrganization.GetOrganizationByID(ctx, &organizationID)
}
