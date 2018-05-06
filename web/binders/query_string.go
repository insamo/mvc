package binders

import (
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

type QueryString struct {
	Filter   string `json:"filter"`
	FilterBy string `json:"filter_by"`

	Sort     []string `json:"sort"`
	Preloads []string `json:"preloads"`

	PerPage int `json:"per_page"`
	Page    int `json:"page"`

	Query map[string][]string `json:"query"`

	CurrPageURL string `json:"curr_page_url"`
}

func NewQueryString(context iris.Context) QueryString {
	var q QueryString
	var query = context.FormValues()
	q.Query = query

	// Filter
	if filter, ok := query["filter"]; ok {
		q.Filter = filter[0]
	}

	// FilterBy
	q.FilterBy = "id"
	if filterBy, ok := query["filterBy"]; ok {
		q.FilterBy = filterBy[0]
	}

	// Sort
	q.Sort = append(q.Sort, "id asc")
	if sorts, ok := query["sort"]; ok {
		// Clear sorts
		q.Sort = q.Sort[:0]
		for _, sort := range sorts {
			s := strings.Split(sort, "|")
			sortQuery := ""
			if len(s) > 1 {
				sortQuery = s[0] + " " + s[1]
			} else {
				sortQuery = s[0]
			}
			q.Sort = append(q.Sort, sortQuery)
		}
	}

	// Preloads relations
	if preloads, ok := query["preloads"]; ok {
		for _, preload := range preloads {
			q.Preloads = append(q.Preloads, preload)
		}
	}

	// Limit
	q.PerPage = 15
	if limit, ok := query["per_page"]; ok {
		if limit[0] != "" {
			q.PerPage, _ = strconv.Atoi(limit[0])
		}
	}

	// Offset
	q.Page = 1
	if offset, ok := query["page"]; ok {
		q.Page, _ = strconv.Atoi(offset[0])
		if q.Page == 0 {
			q.Page = 1
		}
	}

	// Current page url
	q.CurrPageURL = "http://" + context.Host() + context.Request().RequestURI

	return q
}

func (q *QueryString) ProcessQuery(query map[string][]string) {

}
