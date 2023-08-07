package registry

import (
	auth_handler "backend/pkg/adapter/handler/auth"
	organization_handler "backend/pkg/adapter/handler/organization"
	products_handler "backend/pkg/adapter/handler/products"

	"gorm.io/gorm"
)

func AuthServer(db *gorm.DB) auth_handler.AuthController {
	r := NewServer(db)
	return r.NewAuthController()
}

func OrganizationServer(db *gorm.DB) organization_handler.OrganizationController {
	r := NewServer(db)
	return r.NewOrganizationController()
}

func ProductServer(db *gorm.DB) products_handler.ProductController {
	r := NewServer(db)
	return r.NewProductsController()
}
