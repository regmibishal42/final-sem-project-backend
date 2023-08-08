package products_handler

import query_controller "backend/pkg/query/controller"

type ProductRepository struct {
	query_controller.ProductQueryController
}

func NewProductRepository(repo query_controller.ProductQueryController) ProductInterface {
	return ProductRepository{repo}
}
