package registry

import (
	organization_handler "backend/pkg/adapter/handler/organization"
	query_controller "backend/pkg/query/controller"
	query_repository "backend/pkg/query/respository"
)

func (r *Registry) NewOrganizationRegistry() query_controller.OrganizationQueryInterface {
	db := query_controller.OrganizationQueryController{
		TableUser:         query_repository.NewUserQueryRepository(r.db),
		TableProfile:      query_repository.NewProfileQueryRepository(r.db),
		TableOrganization: query_repository.NewOrganizationQueryRepository(r.db),
	}

	return organization_handler.NewOrganizationRepository(db)
}
