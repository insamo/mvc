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
			values = append(values, value)
		}
		s.Params[k] = values
	}

	return s
}


// TODO ParserParams
//func ParseIn() {
//	if v, ok := sqlQuery.Params["city"]; ok {
//		if extras := cityRepo.GetExtrasByUUIDs(txExtra, v); extras != nil {
//			for i := 0; i < len(extras); i++ {
//				sqlQuery.Params["city"][i] = fmt.Sprintf("%d", extras[i].ID)
//			}
//		} else {
//			delete(sqlQuery.Params, "city")
//		}
//	}
//}

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
