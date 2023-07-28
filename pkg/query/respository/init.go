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
