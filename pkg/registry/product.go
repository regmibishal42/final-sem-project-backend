package registry

import (
	products_handler "backend/pkg/adapter/handler/products"
	query_controller "backend/pkg/query/controller"
	query_repository "backend/pkg/query/respository"
)

func (r *Registry) NewProductRegistry() products_handler.ProductInterface {
	db := query_controller.ProductQueryController{
		TableUser:            query_repository.NewUserQueryRepository(r.db),
		TableOrganization:    query_repository.NewOrganizationQueryRepository(r.db),
		TableProduct:         query_repository.NewProductQueryRepository(r.db),
		TableCategory:        query_repository.NewProductCategoryQueryRepository(r.db),
		TableDeletedProducts: query_repository.NewDeletedProductsQueryRepository(r.db),
		TableSales:           query_repository.NewSalesQueryRepository(r.db),
	}
	return products_handler.NewProductRepository(db)
}
