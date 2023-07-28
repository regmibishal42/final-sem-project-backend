package registry

import (
	auth_handler "backend/pkg/adapter/handler/auth"

	"gorm.io/gorm"
)

func AuthServer(db *gorm.DB) auth_handler.AuthController {
	r := NewServer(db)
	return r.NewAuthController()
}
