package query_controller

import (
	"backend/graph/model"
	"context"
)

type ProductQueryInterface interface {
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product, organizationID *string) (*model.Product, error)
	DeleteProduct(ctx context.Context, productID *string, organizationID *string) error

	// Query
	GetProductByID(ctx context.Context, productID *string) (*model.Product, error)
	GetProductsByFilter(ctx context.Context, pageInfo *model.OffsetPageInfo, filter *model.GetProductsByFilterInput, organizationID *string) ([]*model.Product, error)
}
