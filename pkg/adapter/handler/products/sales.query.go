package products_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r ProductRepository) GetSalesByFilter(ctx context.Context, user *model.User, filter *model.FilterSalesInput) (*model.SalesQueryResponse, error) {
	//validate the inputs
	if validationError := filter.Validator(); validationError != nil {
		return &model.SalesQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New(validationError.Message), exception.BAD_REQUEST, nil),
		}, nil
	}
	//check for page info
	var pageInfo model.OffsetPageInfo
	if filter == nil {
		filter = &model.FilterSalesInput{
			Page: &model.OffsetPaginationFilter{},
		}
	} else {
		if filter.Page == nil {
			filter.Page = &model.OffsetPaginationFilter{}
		}
	}
	//get organizationID from user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SalesQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.SalesQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	//get sales
	sales, err := r.TableSales.GetSalesByFilter(ctx, filter, organizationID, &pageInfo)
	if err != nil {
		return &model.SalesQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.SalesQueryResponse{
		Data:     sales,
		PageInfo: &pageInfo,
	}, nil
}

func (r ProductRepository) GetSalesByID(ctx context.Context, user *model.User, input model.GetSalesByIDInput) (*model.SaleQueryResponse, error) {
	//validate the id from input
	if !util.IsValidID(input.SalesID) {
		return &model.SaleQueryResponse{
			Error: exception.QueryErrorHandler(ctx, errors.New("invalid salesID"), exception.BAD_REQUEST, nil),
		}, nil
	}
	//get organizationID from user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SaleQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.SaleQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	sales, err := r.TableSales.GetSalesByID(ctx, &input.SalesID, organizationID)
	if err != nil {
		return &model.SaleQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.SaleQueryResponse{
		Data: sales,
	}, nil
}
