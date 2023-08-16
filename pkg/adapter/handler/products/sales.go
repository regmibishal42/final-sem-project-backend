package products_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (r ProductRepository) CreateSales(ctx context.Context, user *model.User, input *model.CreateSaleInput) (*model.SalesMutationResponse, error) {
	//validate input
	sales, validationError := input.Validator()
	if validationError != nil {
		return &model.SalesMutationResponse{
			Error: validationError,
		}, nil
	}
	//get organization of the user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SalesMutationResponse{
				Error: exception.MutationErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	sales.OrganizationID = *organizationID

	//validate product has enough units
	product, err := r.TableProduct.GetProductByID(ctx, &input.ProductID)
	if err != nil {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	if product.Units < input.Units {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, fmt.Errorf("product only has %v units", product.Units), exception.BAD_REQUEST, nil),
		}, nil
	}
	sales.SoldAt = input.SoldAt
	sales.SoldByID = user.ID
	sales.UnitsSold = input.Units
	//create sales and update the product units
	err = r.TableSales.CreateSales(ctx, sales)
	if err != nil {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	product.Units = product.Units - input.Units
	updatedProduct, err := r.TableProduct.UpdateProduct(ctx, product, organizationID)
	if err != nil {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	sales.Product = updatedProduct
	sales.SoldBy = user
	return &model.SalesMutationResponse{
		ID:   &sales.ID,
		Data: sales,
	}, nil
}

func (r ProductRepository) UpdateSales(ctx context.Context, user *model.User, input *model.UpdateSalesInput) (*model.SalesMutationResponse, error) {
	//validation error
	sales, validationError := input.Validator()
	if validationError != nil {
		return &model.SalesMutationResponse{
			Error: validationError,
		}, nil
	}
	//get organization of the user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SalesMutationResponse{
				Error: exception.MutationErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	// update the sales
	updatedSales, err := r.TableSales.UpdateSales(ctx, sales, organizationID)
	if err != nil {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.SalesMutationResponse{
		ID:   &sales.ID,
		Data: updatedSales,
	}, nil
}

func (r ProductRepository) DeleteSale(ctx context.Context, user *model.User, input *model.DeleteSalesInput) (*model.SalesMutationResponse, error) {
	// Validate ID From The Input
	if !util.IsValidID(input.SalesID) {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, errors.New("invalid salesID"), exception.BAD_REQUEST, nil),
		}, nil
	}
	// Get OrganizationID from the user
	//get organization of the user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SalesMutationResponse{
				Error: exception.MutationErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	err = r.TableSales.DeleteSales(ctx, &input.SalesID, organizationID)
	if err != nil {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}

	return &model.SalesMutationResponse{
		ID: &input.SalesID,
	}, nil
}
