package datasource

import (
	"math"

	"github.com/insamo/mvc/utils"

	"fmt"
	"reflect"
)

// Paginator
type Paginator struct {
	Total       int         `json:"total"`
	PerPage     int64       `json:"per_page"`
	CurrentPage int64       `json:"current_page"`
	LastPage    int64       `json:"last_page"`
	CurrPageUrl string      `json:"curr_page_url"`
	NextPageUrl string      `json:"next_page_url"`
	PrevPageUrl string      `json:"prev_page_url"`
	Params      interface{} `json:"url_params"`
	Data        interface{} `json:"data"`
}

// Paginate results
func NewPaginate(data interface{}, total int, query interface{}) Paginator {
	var (
		p = Paginator{
			Params:      query,
			PerPage:     reflect.ValueOf(query).FieldByName("PerPage").Int(),
			CurrentPage: reflect.ValueOf(query).FieldByName("Page").Int(),
			CurrPageUrl: reflect.ValueOf(query).FieldByName("CurrPageURL").String(),
			Total:       total,
		}
	)

	// LastPage
	p.LastPage = int64(
		math.Ceil(float64(p.Total) / float64(p.PerPage)),
	)

	if p.CurrentPage < p.LastPage {
		p.NextPageUrl = utils.SetUrlQueryString(p.CurrPageUrl, "page", fmt.Sprintf("%d", p.CurrentPage+1))
	} else {
		p.NextPageUrl = ""
	}
	if p.CurrentPage > 1 {
		p.PrevPageUrl = utils.SetUrlQueryString(p.CurrPageUrl, "page", fmt.Sprintf("%d", p.CurrentPage-1))
	} else {
		p.PrevPageUrl = ""
	}

	p.Data = data
	return p
}
