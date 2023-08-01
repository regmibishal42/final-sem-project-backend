package organization_handler

import query_controller "backend/pkg/query/controller"

type OrganizationRepository struct {
	query_controller.OrganizationQueryController
}

func NewOrganizationRepository(repo query_controller.OrganizationQueryController) OrganizationInterface {
	return OrganizationRepository{repo}
}
