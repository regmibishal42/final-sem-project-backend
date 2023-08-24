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

func (r ProductRepository) GetDailySalesStat(ctx context.Context, user *model.User) (*model.DailySalesQueryResponse, error) {
	//get organizationID from user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DailySalesQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.DailySalesQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	//get daily sales data
	data, err := r.TableSales.GetDailySalesStat(ctx, organizationID)
	if err != nil {
		return &model.DailySalesQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}

	return &model.DailySalesQueryResponse{
		Data: data,
	}, nil
}

func (r ProductRepository) GetSalesBreakDownByCategory(ctx context.Context, user *model.User, input *model.SalesBreakDownInput) (*model.SalesBreakDownQueryResponse, error) {
	//get organizationID from user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SalesBreakDownQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.SalesBreakDownQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	//get sales data
	data, err := r.TableSales.GetSalesStatBreakdownByCategory(ctx, *organizationID, input)
	if err != nil {
		return &model.SalesBreakDownQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}

	return &model.SalesBreakDownQueryResponse{
		Data: data,
	}, nil
}

func (r ProductRepository) GetSalesStatByStaff(ctx context.Context, user *model.User, input *model.SalesBreakDownInput) (*model.SalesDataByStaffQueryResponse, error) {
	//get organizationID from user
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SalesDataByStaffQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.SalesDataByStaffQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	//get sales data
	data, err := r.TableSales.GetSalesStatByStaff(ctx, organizationID, input)
	if err != nil {
		return &model.SalesDataByStaffQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.SalesDataByStaffQueryResponse{
		Data: data,
	}, nil
}

func (r ProductRepository) GetDashboardSalesData(ctx context.Context, user *model.User) (*model.DashboardDataQueryResponse, error) {
	organizationID, err := r.TableOrganization.GetOrganizationIDByUser(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DashboardDataQueryResponse{
				Error: exception.QueryErrorHandler(ctx, errors.New("not authorized"), exception.AUTHORIZATION, nil),
			}, nil
		}
		return &model.DashboardDataQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	//get sales data
	data, err := r.TableSales.GetDashboardSalesData(ctx, organizationID)
	if err != nil {
		return &model.DashboardDataQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.DashboardDataQueryResponse{
		Data: data,
	}, nil
}
