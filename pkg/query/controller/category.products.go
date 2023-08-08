package query_controller

import (
	"backend/graph/model"
	"context"
)

type ProductCategoryQueryInterface interface {
	CreateCategory(ctx context.Context, category *model.Category) error
	DeleteCategory(ctx context.Context, categoryID *string) error
	GetCategoryByOrganization(ctx context.Context, organizationID *string) ([]*model.Category, error)
}
