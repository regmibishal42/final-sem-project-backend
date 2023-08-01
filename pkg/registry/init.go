package registry

import (
	auth_handler "backend/pkg/adapter/handler/auth"
	organization_handler "backend/pkg/adapter/handler/organization"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) Registry {
	return Registry{
		db: db,
	}
}

func (r Registry) NewAuthController() auth_handler.AuthController {
	return auth_handler.AuthController{
		AuthDomain: r.NewAuthRegistry(),
	}
}

func (r Registry) NewOrganizationController() organization_handler.OrganizationController {
	return organization_handler.OrganizationController{
		OrganizationDomain: r.NewOrganizationRegistry(),
	}
}
