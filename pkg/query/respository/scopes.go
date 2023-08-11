package query_repository

import (
	"backend/graph/model"
	"math"

	"gorm.io/gorm"
)

func NotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}

func Paginate(value interface{}, page *model.OffsetPageInfo, pagination *model.OffsetPaginationFilter, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	page.TotalRows = int(totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	page.TotalPages = totalPages
	page.Page = pagination.GetPage()

	return func(db *gorm.DB) *gorm.DB {
		if pagination.All != nil {
			if *pagination.All {
				return db.Offset(pagination.GetOffset()).Order(pagination.GetSort())
			}
		}
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
