package products_handler

import (
	"backend/graph/model"
	"context"
)

type ProductController struct {
	ProductDomain ProductInterface
}

type ProductInterface interface {
	// Category
	CreateProductCategory(ctx context.Context, user *model.User, input model.CreateCategoryInput) (*model.CategoryMutationResponse, error)
}
