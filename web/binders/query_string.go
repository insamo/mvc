package binders

import (
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

type QueryString struct {
	Sort []string `json:"sort"`

	PerPage int `json:"per_page"`
	Page    int `json:"page"`

	Query map[string][]string `json:"query"`

	CurrPageURL string `json:"curr_page_url"`
}

func NewQueryString(context iris.Context) QueryString {
	q := QueryString{Query: context.FormValues()}

	if q.Query == nil {
		q.Query = make(map[string][]string)
	}

	// Sort
	q.Sort = append(q.Sort, "id asc")
	if sorts, ok := q.Query["sort"]; ok {
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

	// Limit
	q.PerPage = 15
	if limit, ok := q.Query["per_page"]; ok {
		if limit[0] != "" {
			q.PerPage, _ = strconv.Atoi(limit[0])
		}
	}

	// Offset
	q.Page = 1
	if offset, ok := q.Query["page"]; ok {
		q.Page, _ = strconv.Atoi(offset[0])
		if q.Page == 0 {
			q.Page = 1
		}
	}

	// Current page url
	q.CurrPageURL = "http://" + context.Host() + context.Request().RequestURI

	return q
}
