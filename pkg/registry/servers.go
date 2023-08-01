package registry

import (
	auth_handler "backend/pkg/adapter/handler/auth"
	organization_handler "backend/pkg/adapter/handler/organization"

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
