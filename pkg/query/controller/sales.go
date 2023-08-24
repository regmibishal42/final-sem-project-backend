package query_controller

import (
	"backend/graph/model"
	"context"
)

type SalesQueryInterface interface {
	CreateSales(ctx context.Context, sales *model.Sales) error
	UpdateSales(ctx context.Context, sales *model.Sales, organizationID *string) (*model.Sales, error)
	DeleteSales(ctx context.Context, salesID *string, organizationID *string) error

	GetSalesByFilter(ctx context.Context, filter *model.FilterSalesInput, organizationID *string, pageInfo *model.OffsetPageInfo) ([]*model.Sales, error)
	GetSalesByID(ctx context.Context, salesID *string, organizationID *string) (*model.Sales, error)

	GetSalesStat(ctx context.Context, input *model.SalesStatInput, organizationID *string) (*model.SalesStatData, error)
	GetDailySalesStat(ctx context.Context, organizationID *string) ([]*model.DailySalesData, error)
	GetSalesStatBreakdownByCategory(ctx context.Context, organizationID string, input *model.SalesBreakDownInput) ([]*model.SalesBreakdownData, error)
	GetSalesStatByStaff(ctx context.Context, organizationID *string, input *model.SalesBreakDownInput) ([]*model.SalesDataByStaffs, error)
	GetDashboardSalesData(ctx context.Context, organizationID *string) (*model.DashboardSalesData, error)
}
