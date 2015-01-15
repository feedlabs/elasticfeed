package template

import (
	"strconv"
	"github.com/feedlabs/feedify/context"
)

type ResponseDefinition struct {
	orderby     string
	page        int
	limit       int
}

func (this *ResponseDefinition) GetOrder() string {
	return this.orderby
}

func (this *ResponseDefinition) GetPage() int {
	return this.page
}

func (this *ResponseDefinition) GetLimit() int {
	return this.limit
}

func NewResponseDefinition(input *context.Input) *ResponseDefinition {
	order := input.Request.URL.Query().Get("orderby")
	if order == "" {
		order = "id"
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

	return &ResponseDefinition{order, pageInt, limitInt}
}
