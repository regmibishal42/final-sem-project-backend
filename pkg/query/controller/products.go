package query_controller

import (
	"backend/graph/model"
	"context"
)

type ProductQueryInterface interface {
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
}
