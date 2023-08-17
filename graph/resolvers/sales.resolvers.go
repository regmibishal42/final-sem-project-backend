package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"backend/exception"
	"backend/graph/generated"
	"backend/graph/model"
	"context"
)

// Product is the resolver for the product field.
func (r *salesResolver) Product(ctx context.Context, obj *model.Sales) (*model.Product, error) {
	if obj.Product != nil {
		return obj.Product, nil
	}
	return r.ProductDomain.GetProductDetailsById(ctx, &obj.ProductID)
}

// Organization is the resolver for the organization field.
func (r *salesResolver) Organization(ctx context.Context, obj *model.Sales) (*model.Organization, error) {
	if obj.Organization != nil {
		return obj.Organization, nil
	}
	return r.OrganizationDomain.GetOrganizationDetailsByID(ctx, obj.OrganizationID)
}

// SoldBy is the resolver for the soldBy field.
func (r *salesResolver) SoldBy(ctx context.Context, obj *model.Sales) (*model.User, error) {
	if obj.SoldBy != nil {
		return obj.SoldBy, nil
	}
	return r.AuthDomain.GetUserByID(ctx, &obj.SoldByID)
}

// CreateSales is the resolver for the createSales field.
func (r *salesMutationResolver) CreateSales(ctx context.Context, obj *model.SalesMutation, input model.CreateSaleInput) (*model.SalesMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHENTICATION, nil),
		}, nil
	}
	return r.ProductDomain.CreateSales(ctx, user, &input)
}

// UpdateSales is the resolver for the updateSales field.
func (r *salesMutationResolver) UpdateSales(ctx context.Context, obj *model.SalesMutation, input model.UpdateSalesInput) (*model.SalesMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHENTICATION, nil),
		}, nil
	}
	return r.ProductDomain.UpdateSales(ctx, user, &input)
}

// DeleteSales is the resolver for the deleteSales field.
func (r *salesMutationResolver) DeleteSales(ctx context.Context, obj *model.SalesMutation, input model.DeleteSalesInput) (*model.SalesMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.SalesMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHENTICATION, nil),
		}, nil
	}
	return r.ProductDomain.DeleteSale(ctx, user, &input)
}

// GetSalesByFilter is the resolver for the getSalesByFilter field.
func (r *salesQueryResolver) GetSalesByFilter(ctx context.Context, obj *model.SalesQuery, input model.FilterSalesInput) (*model.SalesQueryResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.SalesQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.AUTHENTICATION, nil),
		}, nil
	}
	return r.ProductDomain.GetSalesByFilter(ctx, user, &input)
}

// GetSaleByID is the resolver for the getSaleByID field.
func (r *salesQueryResolver) GetSaleByID(ctx context.Context, obj *model.SalesQuery, input model.GetSalesByIDInput) (*model.SaleQueryResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.SaleQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.AUTHENTICATION, nil),
		}, nil
	}
	return r.ProductDomain.GetSalesByID(ctx, user, input)
}

// Sales returns generated.SalesResolver implementation.
func (r *Resolver) Sales() generated.SalesResolver { return &salesResolver{r} }

// SalesMutation returns generated.SalesMutationResolver implementation.
func (r *Resolver) SalesMutation() generated.SalesMutationResolver { return &salesMutationResolver{r} }

// SalesQuery returns generated.SalesQueryResolver implementation.
func (r *Resolver) SalesQuery() generated.SalesQueryResolver { return &salesQueryResolver{r} }

type salesResolver struct{ *Resolver }
type salesMutationResolver struct{ *Resolver }
type salesQueryResolver struct{ *Resolver }
