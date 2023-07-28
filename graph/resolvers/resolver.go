package resolvers

import auth_handler "backend/pkg/adapter/handler/auth"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	auth_handler.AuthController
}

func NewResolver(authController auth_handler.AuthController) *Resolver {
	resolver := &Resolver{
		authController,
	}
	return resolver
}
