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

func (r OrganizationRepository) GetStaffByID(ctx context.Context, user *model.User, input *model.GetStaffInput) (*model.StaffQueryResponse, error) {
	if !util.IsValidID(input.StaffID) {
		return &model.StaffQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("invalid StaffID"), exception.BAD_REQUEST, nil),
		}, nil
	}
	//cannot view other staff details
	if user.UserType == model.UserTypeStaff {
		if user.ID != input.StaffID {
			return &model.StaffQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("not authorized for this task"), exception.AUTHORIZATION, nil),
			}, nil
		}
	}
	// Get Staff
	staff, err := r.TableStaff.GetStaffByID(ctx, &input.StaffID)
	if err != nil {
		return &model.StaffQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.BAD_REQUEST, nil),
		}, nil
	}
	return &model.StaffQueryResponse{
		Data: staff,
	}, nil
}

func (r OrganizationRepository) GetStaffsByOrganization(ctx context.Context, user *model.User, input *model.GetOrganizationStaffsInput) (*model.StaffsQueryResponse, error) {
	//check validity of organizationID
	if !util.IsValidID(input.OrganizationID) {
		return &model.StaffsQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("invalid StaffID"), exception.BAD_REQUEST, nil),
		}, nil
	}
	if user.UserType == model.UserTypeStaff {
		return &model.StaffsQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("not authorized for this task"), exception.AUTHORIZATION, nil),
		}, nil
	}
	//get all staffs of an organization
	staffs, err := r.TableStaff.GetStaffsByOrganization(ctx, &input.OrganizationID)
	if err != nil {
		return &model.StaffsQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.BAD_REQUEST, nil),
		}, nil
	}
	return &model.StaffsQueryResponse{
		Data: staffs,
	}, nil
}
