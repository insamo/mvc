package repositories

import (
	"github.com/insamo/mvc/web/binders"
	"github.com/jinzhu/gorm"
)

func Prepare(query binders.QueryString, db *gorm.DB) *gorm.DB {
	db = db.Select("*, count(*) OVER() AS full_count")

	// Order
	if len(query.Sort) > 0 {
		for _, sort := range query.Sort {
			db = db.Order(sort)
		}
	}

	// Limit
	if query.PerPage > 0 {
		db = db.Limit(query.PerPage)
	}

	// Offset
	db = db.Offset((query.Page - 1) * query.PerPage)

	return db
}
