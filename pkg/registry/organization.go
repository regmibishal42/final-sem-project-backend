package registry

import (
	organization_handler "backend/pkg/adapter/handler/organization"
	query_controller "backend/pkg/query/controller"
	query_repository "backend/pkg/query/respository"
)

func (r *Registry) NewOrganizationRegistry() organization_handler.OrganizationInterface {
	db := query_controller.OrganizationQueryController{
		TableUser:         query_repository.NewUserQueryRepository(r.db),
		TableProfile:      query_repository.NewProfileQueryRepository(r.db),
		TableOrganization: query_repository.NewOrganizationQueryRepository(r.db),
		TableStaff:        query_repository.NewStaffQueryRepository(r.db),
		TableProduct:      query_repository.NewProductQueryRepository(r.db),
	}

	return organization_handler.NewOrganizationRepository(db)
}
