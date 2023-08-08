package products_handler

import (
	"backend/exception"
	"backend/graph/model"
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
