package registry

import (
	auth_handler "backend/pkg/adapter/handler/auth"
	query_controller "backend/pkg/query/controller"
	query_repository "backend/pkg/query/respository"
)

func (r *Registry) NewAuthRegistry() auth_handler.AuthInterface {
	db := query_controller.AuthQueryController{
		TableUser: query_repository.NewUserQueryRepository(r.db),
	}

	return auth_handler.NewAuthRepository(db)
}
