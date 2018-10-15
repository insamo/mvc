package binders

import (
	"strconv"

	"github.com/kataras/iris"
)

type QueryString struct {
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

	// Remove empty query params
	for k, v := range q.Query {
		for i := 0; i < len(v); i++ {
			if v[i] == "" {
				v = append(v[:i], v[i+1:]...)
			}
		}
		q.Query[k] = v
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
	q.Page = (q.Page - 1) * q.PerPage

	// Current page url
	q.CurrPageURL = "http://" + context.Host() + context.Request().RequestURI

	return q
}
