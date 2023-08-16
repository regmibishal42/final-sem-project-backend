package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"backend/graph/generated"
	"backend/graph/model"
	"context"
	"fmt"
)

// Product is the resolver for the product field.
func (r *salesResolver) Product(ctx context.Context, obj *model.Sales) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: Product - product"))
}

// Organization is the resolver for the organization field.
func (r *salesResolver) Organization(ctx context.Context, obj *model.Sales) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// SoldBy is the resolver for the soldBy field.
func (r *salesResolver) SoldBy(ctx context.Context, obj *model.Sales) (*model.User, error) {
	panic(fmt.Errorf("not implemented: SoldBy - soldBy"))
}

// CreateSales is the resolver for the createSales field.
func (r *salesMutationResolver) CreateSales(ctx context.Context, obj *model.SalesMutation, input model.CreateSaleInput) (*model.SalesMutationResponse, error) {
	panic(fmt.Errorf("not implemented: CreateSales - createSales"))
}

// UpdateSales is the resolver for the updateSales field.
func (r *salesMutationResolver) UpdateSales(ctx context.Context, obj *model.SalesMutation, input model.UpdateSalesInput) (*model.SalesMutationResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateSales - updateSales"))
}

// DeleteSales is the resolver for the deleteSales field.
func (r *salesMutationResolver) DeleteSales(ctx context.Context, obj *model.SalesMutation, input model.DeleteSalesInput) (*model.SalesMutationResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteSales - deleteSales"))
}

// GetSalesByFilter is the resolver for the getSalesByFilter field.
func (r *salesQueryResolver) GetSalesByFilter(ctx context.Context, obj *model.SalesQuery, input model.FilterSalesInput) (*model.SalesQueryResponse, error) {
	panic(fmt.Errorf("not implemented: GetSalesByFilter - getSalesByFilter"))
}

// GetSaleByID is the resolver for the getSaleByID field.
func (r *salesQueryResolver) GetSaleByID(ctx context.Context, obj *model.SalesQuery, input model.GetSalesByIDInput) (*model.SaleQueryResponse, error) {
	panic(fmt.Errorf("not implemented: GetSaleByID - getSaleByID"))
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
