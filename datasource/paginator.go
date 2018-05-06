package datasource

import (
	"math"
	"strconv"

	"bitbucket.org/insamo/mvc/utils"

	"bitbucket.org/insamo/mvc/web/binders"
)

// Paginator type for pagination
type Paginator struct {
	Total       int                 `json:"total"`
	PerPage     int                 `json:"per_page"`
	CurrentPage int                 `json:"current_page"`
	LastPage    int                 `json:"last_page"`
	CurrPageUrl string              `json:"curr_page_url"`
	NextPageUrl string              `json:"next_page_url"`
	PrevPageUrl string              `json:"prev_page_url"`
	From        int                 `json:"from"`
	To          int                 `json:"to"`
	Params      binders.QueryString `json:"url_params"`
	Data        interface{}         `json:"data"`
}

// Paginate results
func NewPaginate(data interface{}, total int, query binders.QueryString) Paginator {
	var p Paginator

	p.Params = query

	p.Total = total
	p.PerPage = query.PerPage
	// LastPage
	d := float64(p.Total) / float64(p.PerPage)
	p.LastPage = int(math.Ceil(d))

	// Offset
	p.CurrentPage = query.Page
	p.CurrPageUrl = query.CurrPageURL

	if p.CurrentPage < p.LastPage {
		p.NextPageUrl = utils.SetUrlQueryString(p.CurrPageUrl, "page", strconv.Itoa(p.CurrentPage+1))
	} else {
		p.NextPageUrl = ""
	}
	if p.CurrentPage > 1 {
		p.PrevPageUrl = utils.SetUrlQueryString(p.CurrPageUrl, "page", strconv.Itoa(p.CurrentPage-1))
	} else {
		p.PrevPageUrl = ""
	}

	p.From = 0 //TODO
	p.To = 0   //TODO

	p.Data = data
	return p
}
