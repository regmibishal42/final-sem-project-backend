package registry

import (
	auth_handler "backend/pkg/adapter/handler/auth"

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
