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
	UpdateProduct(ctx context.Context, user *model.User, input *model.UpdateProductInput) (*model.ProductMutationResponse, error)
	DeleteProduct(ctx context.Context, user *model.User, productID *string) (*model.ProductMutationResponse, error)

	GetProductByID(ctx context.Context, user *model.User, productID *string) (*model.ProductQueryResponse, error)
	GetProductDetailsById(ctx context.Context, productID *string) (*model.Product, error)
	GetProductsByFilter(ctx context.Context, user *model.User, filter *model.GetProductsByFilterInput) (*model.ProductsQueryResponse, error)

	//sales
	CreateSales(ctx context.Context, user *model.User, input *model.CreateSaleInput) (*model.SalesMutationResponse, error)
}
