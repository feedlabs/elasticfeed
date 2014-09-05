package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/feedlabs/feedify/lib/feedify/service"
)

type ResponseInfo struct {
}

type DefaultController struct {
	beego.Controller
}

func (this *DefaultController) Get() {
	this.Data["json"] = map[string]string{"succes": "ok"}
	this.ServeJson()
}

func GetResponseFormat(input *context.BeegoInput) string {
	format := "json"
	parts := strings.Split(input.Uri(), ".")
	if (len(parts) > 1) {
		format = parts[len(parts) - 1]
	}
	return format
}

func GetRequestParam(input *context.BeegoInput, param string) string {
	return input.Query(param)
}

func init() {
	neo4j := service.NewNeo4j()
	neo4j.Connect()

	memcache := service.NewMemcache()
	memcache.Connect()
}
