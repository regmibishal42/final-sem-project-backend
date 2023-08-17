package query_controller

import (
	"backend/graph/model"
	"context"
)

type SalesQueryInterface interface {
	CreateSales(ctx context.Context, sales *model.Sales) error
	UpdateSales(ctx context.Context, sales *model.Sales, organizationID *string) (*model.Sales, error)
	DeleteSales(ctx context.Context, salesID *string, organizationID *string) error

	GetSalesByFilter(ctx context.Context, filter *model.FilterSalesInput, organizationID *string) ([]*model.Sales, error)
}
