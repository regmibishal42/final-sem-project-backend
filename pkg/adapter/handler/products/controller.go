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
	DeleteCategory(ctx context.Context, user *model.User, input model.DeleteCategoryInput) (*model.CategoryMutationResponse, error)
	GetCategoryByOrganization(ctx context.Context, user *model.User) (*model.CategoryQueryResponse, error)
	GetCategoryByID(ctx context.Context, categoryID *string) (*model.Category, error)

	// Products
	CreateProduct(ctx context.Context, user *model.User, input *model.CreateProductInput) (*model.ProductMutationResponse, error)
}
