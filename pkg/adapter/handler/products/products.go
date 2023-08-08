package products_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
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

//update products details --Add User Verification Later
func (r ProductRepository) UpdateProduct(ctx context.Context, user *model.User, input *model.UpdateProductInput) (*model.ProductMutationResponse, error) {
	product, validationError := input.Validator()
	if validationError != nil {
		return &model.ProductMutationResponse{
			Error: validationError,
		}, nil
	}

	//update the product
	updatedProduct, err := r.TableProduct.UpdateProduct(ctx, product)
	if err != nil {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.ProductMutationResponse{
		ID:   &product.ID,
		Data: updatedProduct,
	}, nil
}

//delete product
func (r ProductRepository) DeleteProduct(ctx context.Context, user *model.User, productID *string) (*model.ProductMutationResponse, error) {
	//validate the productID
	if !util.IsValidID(*productID) {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("invalid productID"), exception.BAD_REQUEST, nil),
		}, nil
	}
	//Delete the product
	err := r.TableProduct.DeleteProduct(ctx, productID)
	if err != nil {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}

	return &model.ProductMutationResponse{
		ID: productID,
	}, nil
}
