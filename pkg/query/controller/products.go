package query_controller

import (
	"backend/graph/model"
	"context"
)

type ProductQueryInterface interface {
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, productID *string) error

	// Query
	GetProductByID(ctx context.Context, productID *string) (*model.Product, error)
	GetProductsByFilter(ctx context.Context, filter *model.GetProductsByFilterInput, organizationID *string) ([]*model.Product, error)
}
