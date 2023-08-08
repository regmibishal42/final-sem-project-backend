package query_controller

import (
	"backend/graph/model"
	"context"
)

type ProductCategoryQueryInterface interface {
	CreateCategory(ctx context.Context, category *model.Category) error
}
