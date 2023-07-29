package query_controller

type AuthQueryController struct {
	//tables eg TableUser
	TableUser    UserQueryInterface
	TableProfile ProfileQueryInterface
}
