package products_handler

import (
	"backend/exception"
	"backend/graph/model"
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
	sales, err := r.TableSales.GetSalesByFilter(ctx, filter, organizationID)
	if err != nil {
		return &model.SalesQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.SalesQueryResponse{
		Data: sales,
	}, nil
}
