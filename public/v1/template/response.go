package template

import (
	"strconv"
	"github.com/feedlabs/feedify/context"
)

type ResponseDefinition struct {
	orderby     string
	orderdir    string
	page        int
	limit       int
}

func (this *ResponseDefinition) GetOrderBy() string {
	return this.orderby
}

func (this *ResponseDefinition) GetOrderDir() string {
	return this.orderdir
}

func (this *ResponseDefinition) GetPage() int {
	return this.page
}

func (this *ResponseDefinition) GetLimit() int {
	return this.limit
}

func NewResponseDefinition(input *context.Input) *ResponseDefinition {
	orderby := input.Request.URL.Query().Get("orderby")
	if orderby == "" {
		orderby = "id"
	}

	orderdir := input.Request.URL.Query().Get("orderdir")
	if orderdir == "" {
		orderdir = "asc"
	}

	page := input.Request.URL.Query().Get("page")
	if page == "" {
		page = "0"
	}

	limit := input.Request.URL.Query().Get("limit")
	if limit == "" {
		limit = "100"
	}

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	return &ResponseDefinition{orderby, orderdir, pageInt, limitInt}
}
