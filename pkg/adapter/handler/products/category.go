package products_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
)

func (r ProductRepository) CreateProductCategory(ctx context.Context, user *model.User, input model.CreateCategoryInput) (*model.CategoryMutationResponse, error) {
	if validationError := input.Validator(); validationError != nil {
		return &model.CategoryMutationResponse{
			Error: validationError,
		}, nil
	}
	id, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	if id == nil {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("not authorized for this task"), exception.AUTHORIZATION, nil),
		}, nil
	}
	//create Category
	category := model.Category{
		Name:           input.Name,
		OrganizationID: *id,
	}
	err = r.TableCategory.CreateCategory(ctx, &category)
	if err != nil {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.CategoryMutationResponse{
		Data: &category,
	}, nil

}

func (r ProductRepository) DeleteCategory(ctx context.Context, user *model.User, input model.DeleteCategoryInput) (*model.CategoryMutationResponse, error) {
	//validate the id from input
	if !util.IsValidID(input.CategoryID) {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("invalid CategoryID"), exception.BAD_REQUEST, nil),
		}, nil
	}
	//Get OrganizationID from userID
	id, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	if id == nil {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("not authorized for this task"), exception.AUTHORIZATION, nil),
		}, nil
	}
	//Soft delete category
	err = r.TableCategory.DeleteCategory(ctx, &input.CategoryID)
	if err != nil {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.CategoryMutationResponse{
		ID: &input.CategoryID,
	}, nil
}

func (r ProductRepository) GetCategoryByOrganization(ctx context.Context, user *model.User) (*model.CategoryQueryResponse, error) {
	//Get OrganizationID from userID
	id, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		return &model.CategoryQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	if id == nil {
		return &model.CategoryQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("not authorized for this task"), exception.AUTHORIZATION, nil),
		}, nil
	}
	//get category
	categories, err := r.TableCategory.GetCategoryByOrganization(ctx, id)
	if err != nil {
		return &model.CategoryQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.CategoryQueryResponse{
		Data: categories,
	}, nil
}

func (r ProductRepository) GetCategoryByID(ctx context.Context, categoryID *string) (*model.Category, error) {
	return r.TableCategory.GetCategoryByID(ctx, categoryID)
}
