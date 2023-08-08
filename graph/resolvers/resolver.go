package resolvers

import (
	auth_handler "backend/pkg/adapter/handler/auth"
	organization_handler "backend/pkg/adapter/handler/organization"
	products_handler "backend/pkg/adapter/handler/products"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	auth_handler.AuthController
	organization_handler.OrganizationController
	products_handler.ProductController
}

func NewResolver(authController auth_handler.AuthController,
	organizationController organization_handler.OrganizationController,
	productController products_handler.ProductController) *Resolver {
	resolver := &Resolver{
		authController,
		organizationController,
		productController,
	}
	return resolver
}
