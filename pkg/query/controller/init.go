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
	TableProduct      ProductQueryInterface
}

type ProductQueryController struct {
	TableProduct         ProductQueryInterface
	TableOrganization    OrganizationQueryInterface
	TableUser            UserQueryInterface
	TableCategory        ProductCategoryQueryInterface
	TableDeletedProducts DeletedProductsQueryInterface
}
