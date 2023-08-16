package products_handler

import (
	"backend/graph/model"
	"context"
)

func (r ProductRepository) CreateSales(ctx context.Context, user *model.User, input *model.CreateSaleInput) (*model.SalesMutationResponse, error) {

	return &model.SalesMutationResponse{
		Data: nil,
	}, nil
}
