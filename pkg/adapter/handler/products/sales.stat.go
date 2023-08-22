package products_handler

import (
	"backend/exception"
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r ProductRepository) GetSalesStatOverview(ctx context.Context, user *model.User, input *model.SalesStatInput) (*model.SalesStatQueryResponse, error) {
	//validate input if exist
	if input != nil {
		if input.CategoryID != nil && util.IsValidID(*input.CategoryID) {
			return &model.SalesStatQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("invalid categoryID"), exception.BAD_REQUEST, nil),
			}, nil
		}
	}
	//get organizationID from user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SalesStatQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.SalesStatQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	//get sales stat data
	salesStat, err := r.TableSales.GetSalesStat(ctx, input, organizationID)
	if err != nil {
		return &model.SalesStatQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.SalesStatQueryResponse{
		Data: salesStat,
	}, nil
}
