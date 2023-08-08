package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"backend/exception"
	"backend/graph/generated"
	"backend/graph/model"
	"context"
	"fmt"
)

// Category is the resolver for the category field.
func (r *productResolver) Category(ctx context.Context, obj *model.Product) (*model.Category, error) {
	if obj.Category != nil {
		return obj.Category, nil
	}
	return r.ProductDomain.GetCategoryByID(ctx, &obj.CategoryID)
}

// Organization is the resolver for the organization field.
func (r *productResolver) Organization(ctx context.Context, obj *model.Product) (*model.Organization, error) {
	if obj.Organization != nil {
		return obj.Organization, nil
	}
	return r.OrganizationDomain.GetOrganizationDetailsByID(ctx, obj.OrganizationID)
}

// Category is the resolver for the category field.
func (r *productMutationResolver) Category(ctx context.Context, obj *model.ProductMutation) (*model.CategoryMutation, error) {
	return &model.CategoryMutation{}, nil
}

// CreateProduct is the resolver for the createProduct field.
func (r *productMutationResolver) CreateProduct(ctx context.Context, obj *model.ProductMutation, input model.CreateProductInput) (*model.ProductMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	return r.ProductDomain.CreateProduct(ctx, user, &input)
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *productMutationResolver) UpdateProduct(ctx context.Context, obj *model.ProductMutation, input model.UpdateProductInput) (*model.ProductMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	return r.ProductDomain.UpdateProduct(ctx, user, &input)
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *productMutationResolver) DeleteProduct(ctx context.Context, obj *model.ProductMutation, input model.DeleteProductInput) (*model.ProductMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.ProductMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	return r.ProductDomain.DeleteProduct(ctx, user, &input.ProductID)
}

// Category is the resolver for the category field.
func (r *productQueryResolver) Category(ctx context.Context, obj *model.ProductQuery) (*model.CategoryQuery, error) {
	return &model.CategoryQuery{}, nil
}

// GetProductsByFilter is the resolver for the getProductsByFilter field.
func (r *productQueryResolver) GetProductsByFilter(ctx context.Context, obj *model.ProductQuery, input model.GetProductsByFilterInput) (*model.ProductsQueryResponse, error) {
	panic(fmt.Errorf("not implemented: GetProductsByFilter - getProductsByFilter"))
}

// GetProductByID is the resolver for the getProductByID field.
func (r *productQueryResolver) GetProductByID(ctx context.Context, obj *model.ProductQuery, input model.GetProductByIDInput) (*model.ProductQueryResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.ProductQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	return r.ProductDomain.GetProductByID(ctx, user, &input.ProductID)
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// ProductMutation returns generated.ProductMutationResolver implementation.
func (r *Resolver) ProductMutation() generated.ProductMutationResolver {
	return &productMutationResolver{r}
}

// ProductQuery returns generated.ProductQueryResolver implementation.
func (r *Resolver) ProductQuery() generated.ProductQueryResolver { return &productQueryResolver{r} }

type productResolver struct{ *Resolver }
type productMutationResolver struct{ *Resolver }
type productQueryResolver struct{ *Resolver }
