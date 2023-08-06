package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"backend/exception"
	"backend/graph/generated"
	"backend/graph/model"
	"context"
	"fmt"
)

// Staff is the resolver for the staff field.
func (r *staffResolver) Staff(ctx context.Context, obj *model.Staff) (*model.User, error) {
	return r.AuthDomain.GetUserByID(ctx, &obj.StaffID)
}

// Organization is the resolver for the Organization field.
func (r *staffResolver) Organization(ctx context.Context, obj *model.Staff) (*model.Organization, error) {
	return r.OrganizationDomain.GetOrganizationDetailsByID(ctx, obj.OrganizationID)
}

// CreateStaff is the resolver for the createStaff field.
func (r *staffMutationResolver) CreateStaff(ctx context.Context, obj *model.StaffMutation, input model.CreateStaffInput) (*model.StaffMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.StaffMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	return r.OrganizationDomain.CreateStaff(ctx, user, &input)
}

// UpdateStaff is the resolver for the updateStaff field.
func (r *staffMutationResolver) UpdateStaff(ctx context.Context, obj *model.StaffMutation, input model.UpdateStaffInput) (*model.StaffMutationResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateStaff - updateStaff"))
}

// GetStaffByOrganization is the resolver for the getStaffByOrganization field.
func (r *staffQueryResolver) GetStaffByOrganization(ctx context.Context, obj *model.StaffQuery, input model.GetOrganizationStaffsInput) (*model.StaffsQueryResponse, error) {
	panic(fmt.Errorf("not implemented: GetStaffByOrganization - getStaffByOrganization"))
}

// GetStaffByID is the resolver for the getStaffByID field.
func (r *staffQueryResolver) GetStaffByID(ctx context.Context, obj *model.StaffQuery, input model.GetStaffInput) (*model.StaffQueryResponse, error) {
	panic(fmt.Errorf("not implemented: GetStaffByID - getStaffByID"))
}

// Staff returns generated.StaffResolver implementation.
func (r *Resolver) Staff() generated.StaffResolver { return &staffResolver{r} }

// StaffMutation returns generated.StaffMutationResolver implementation.
func (r *Resolver) StaffMutation() generated.StaffMutationResolver { return &staffMutationResolver{r} }

// StaffQuery returns generated.StaffQueryResolver implementation.
func (r *Resolver) StaffQuery() generated.StaffQueryResolver { return &staffQueryResolver{r} }

type staffResolver struct{ *Resolver }
type staffMutationResolver struct{ *Resolver }
type staffQueryResolver struct{ *Resolver }
