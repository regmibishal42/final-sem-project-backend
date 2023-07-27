package db

import (
	"backend/graph/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
	//migration models
	)
	DbExceptionHandle(err)
}
