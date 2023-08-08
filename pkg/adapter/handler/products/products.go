package products_handler

import (
	"backend/exception"
	"backend/graph/model"
	"context"
	"errors"
)

func (r ProductRepository) CreateProduct(ctx context.Context, user *model.User, input *model.CreateProductInput) (*model.ProductMutationResponse, error) {
	//validate input
	product, validationError := input.Validator()
	if validationError != nil {
		return &model.ProductMutationResponse{
			Error: validationError,
		}, nil
	}
	//check for user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	if organizationID == nil {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("not authorized for this task"), exception.AUTHORIZATION, nil),
		}, nil
	}
	//add organizationID to product
	product.OrganizationID = *organizationID
	//create product
	err = r.TableProduct.CreateProduct(ctx, product)
	if err != nil {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.ProductMutationResponse{
		ID:   &product.ID,
		Data: product,
	}, nil
}
