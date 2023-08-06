package query_controller

type AuthQueryController struct {
	TableUser    UserQueryInterface
	TableProfile ProfileQueryInterface
	TableOtp     OtpQueryInterface
}

type OrganizationQueryController struct {
	TableUser         UserQueryInterface
	TableProfile      ProfileQueryInterface
	TableOrganization OrganizationQueryInterface
	TableStaff        StaffQueryInterface
}
