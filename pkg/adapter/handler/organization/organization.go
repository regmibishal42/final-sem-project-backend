package organization_handler

import (
	"backend/exception"
	"backend/graph/model"
	"context"
	"errors"
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

	organization.CreatedBy = user
	organization.CreatedByID = user.ID

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
