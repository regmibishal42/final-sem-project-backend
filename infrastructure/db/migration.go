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
	//migration models
	)
	DbExceptionHandle(err)
}
