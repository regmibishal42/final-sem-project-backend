package db

import (
	"backend/graph/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Profile{},
		&model.Otp{},
		&model.Organization{},
		&model.Staff{},
		&model.DeletedProducts{},
		&model.Product{},
		&model.Category{},
		&model.Sales{},
	//migration models
	)
	DbExceptionHandle(err)
}
