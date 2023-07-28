package auth_handler

import query_controller "backend/pkg/query/controller"

type AuthRepository struct {
	query_controller.AuthQueryController
}

func NewAuthRepository(repo query_controller.AuthQueryController) AuthInterface {
	return AuthRepository{repo}
}
