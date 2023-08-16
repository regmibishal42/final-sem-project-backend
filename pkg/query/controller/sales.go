package query_controller

import (
	"backend/graph/model"
	"context"
)

type SalesQueryInterface interface {
	CreateSales(ctx context.Context, sales *model.Sales) error
	UpdateSales(ctx context.Context, sales *model.Sales) (*model.Sales, error)
}