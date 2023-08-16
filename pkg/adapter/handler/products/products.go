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

//get product by id
func (r ProductRepository) GetProductByID(ctx context.Context, user *model.User, productID *string) (*model.ProductQueryResponse, error) {
	//validate productID
	if !util.IsValidID(*productID) {
		return &model.ProductQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("invalid productID"), exception.BAD_REQUEST, nil),
		}, nil
	}
	//validate the user -> permission
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		return &model.ProductQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	if organizationID == nil {
		return &model.ProductQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("not authorized for this task"), exception.AUTHORIZATION, nil),
		}, nil
	}
	//get the product
	product, err := r.TableProduct.GetProductByID(ctx, productID)
	if err != nil {
		return &model.ProductQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.ProductQueryResponse{
		Data: product,
	}, nil
}

func (r ProductRepository) GetProductsByFilter(ctx context.Context, user *model.User, filter *model.GetProductsByFilterInput) (*model.ProductsQueryResponse, error) {
	//validate category id, if exist
	if filter != nil && filter.Params != nil && filter.Params.CategoryID != nil && !util.IsValidID(*filter.Params.CategoryID) {
		return &model.ProductsQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("invalid categoryID"), exception.BAD_REQUEST, nil),
		}, nil
	}
	//pagination
	var pageInfo model.OffsetPageInfo
	if filter == nil {
		filter = &model.GetProductsByFilterInput{
			Page: &model.OffsetPaginationFilter{},
		}
	} else {
		if filter.Page == nil {
			filter.Page = &model.OffsetPaginationFilter{}
		}
	}

	//validate the user -> permission
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		return &model.ProductsQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	if organizationID == nil {
		return &model.ProductsQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("not authorized for this task"), exception.AUTHORIZATION, nil),
		}, nil
	}
	// get the products
	products, err := r.TableProduct.GetProductsByFilter(ctx, &pageInfo, filter, organizationID)
	if err != nil {
		return &model.ProductsQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.ProductsQueryResponse{
		Data:     products,
		PageInfo: &pageInfo,
	}, nil
}

func (r ProductRepository) GetProductDetailsById(ctx context.Context, productID *string) (*model.Product, error) {
	return r.TableProduct.GetProductByID(ctx, productID)
}
