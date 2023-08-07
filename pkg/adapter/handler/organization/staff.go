package organization_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
)

func (r OrganizationRepository) CreateStaff(ctx context.Context, user *model.User, input *model.CreateStaffInput) (*model.StaffMutationResponse, error) {
	if validationError := input.Validator(); validationError != nil {
		return &model.StaffMutationResponse{
			Error: validationError,
		}, nil
	}
	//create staff
	staff := model.Staff{
		OrganizationID: input.OrganizationID,
		JoinedOn:       input.JoinedOn,
		Post:           input.Post,
		Salary:         input.Salary,
		IsAuthorized:   input.IsAuthorized,
		Staff: &model.User{
			Email:      input.Email,
			UserType:   model.UserTypeStaff,
			IsVerified: false,
			Status:     model.UserStatusActive,
			Password:   util.GetDefaultStaffPassword(),
			Profile: &model.Profile{
				FirstName:     input.FirstName,
				LastName:      input.LastName,
				ContactNumber: input.ContactNumber,
				Address:       (*model.Address)(input.Address),
			},
		},
	}
	if input.IsAuthorized != nil && util.Deref(input.IsAuthorized) {
		staffName := input.FirstName + " " + input.LastName
		go util.SendStaffAccountCreationEmail(input.Email, staffName)
	}

	err := r.TableStaff.CreateStaff(ctx, &staff)
	if err != nil {
		return &model.StaffMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.StaffMutationResponse{
		Data: &staff,
	}, nil
}

func (r OrganizationRepository) UpdateStaffDetails(ctx context.Context, user *model.User, input *model.UpdateStaffInput) (*model.StaffMutationResponse, error) {
	staff, validationError := input.Validator()
	if validationError != nil {
		return &model.StaffMutationResponse{
			Error: validationError,
		}, nil
	}
	if user.UserType != model.UserTypeAdmin {
		return &model.StaffMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
		}, nil
	}
	err := r.TableStaff.UpdateStaff(ctx, staff)
	if err != nil {
		return &model.StaffMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.StaffMutationResponse{
		Data: staff,
	}, nil
}
