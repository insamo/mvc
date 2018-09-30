package repositories

import (
	"github.com/insamo/mvc/web/binders"
	"github.com/jinzhu/gorm"
)

type SqlQuery struct {
	Limit  int
	Offset int
	Params map[string][]string
}

func NewSqlQuery(queryString binders.QueryString) SqlQuery {

	s := SqlQuery{
		Limit:  queryString.PerPage,
		Offset: (queryString.Page - 1) * queryString.PerPage,
	}

	// Copy map oO
	s.Params = make(map[string][]string)
	for k, v := range queryString.Query {
		values := []string{}
		for _, value := range v {
			if value != "" {
				values = append(values, value)
			}
		}
		s.Params[k] = values
	}

	return s
}

func Prepare(query SqlQuery, db *gorm.DB) *gorm.DB {
	db = db.Select("*, count(*) OVER() AS full_count")

	// Limit
	if query.Limit > 0 {
		db = db.Limit(query.Limit)
	}

	// Offset
	db = db.Offset(query.Offset)

	return db
}
