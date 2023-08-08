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

// CreateCategory is the resolver for the createCategory field.
func (r *categoryMutationResolver) CreateCategory(ctx context.Context, obj *model.CategoryMutation, input model.CreateCategoryInput) (*model.CategoryMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	return r.ProductDomain.CreateProductCategory(ctx, user, input)
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *categoryMutationResolver) DeleteCategory(ctx context.Context, obj *model.CategoryMutation, input model.DeleteCategoryInput) (*model.CategoryMutationResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.CategoryMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	panic(fmt.Errorf("not implemented: DeleteCategory - deleteCategory"))
}

// GetAllCategory is the resolver for the getAllCategory field.
func (r *categoryQueryResolver) GetAllCategory(ctx context.Context, obj *model.CategoryQuery, input model.GetCategoriesInput) (*model.CategoryQueryResponse, error) {
	user := UserForContext(ctx)
	err := CheckLoggedIn(user)
	if err != nil {
		return &model.CategoryQueryResponse{
			Error: exception.QueryErrorHandler(ctx, err, exception.AUTHORIZATION, nil),
		}, nil
	}
	panic(fmt.Errorf("not implemented: GetAllCategory - getAllCategory"))
}

// CategoryMutation returns generated.CategoryMutationResolver implementation.
func (r *Resolver) CategoryMutation() generated.CategoryMutationResolver {
	return &categoryMutationResolver{r}
}

// CategoryQuery returns generated.CategoryQueryResolver implementation.
func (r *Resolver) CategoryQuery() generated.CategoryQueryResolver { return &categoryQueryResolver{r} }

type categoryMutationResolver struct{ *Resolver }
type categoryQueryResolver struct{ *Resolver }
