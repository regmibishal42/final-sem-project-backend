package query_repository

import (
	query_controller "backend/pkg/query/controller"

	"gorm.io/gorm"
)

type QueryRepository struct {
	db *gorm.DB
}

func NewUserQueryRepository(db *gorm.DB) query_controller.UserQueryInterface {
	return QueryRepository{
		db: db,
	}
}

func NewProfileQueryRepository(db *gorm.DB) query_controller.ProfileQueryInterface {
	return QueryRepository{
		db: db,
	}
}

func NewOtpQueryRepository(db *gorm.DB) query_controller.OtpQueryInterface {
	return QueryRepository{
		db: db,
	}
}

func NewOrganizationQueryRepository(db *gorm.DB) query_controller.OrganizationQueryInterface {
	return QueryRepository{
		db: db,
	}
}

func NewStaffQueryRepository(db *gorm.DB) query_controller.StaffQueryInterface {
	return QueryRepository{
		db: db,
	}
}
