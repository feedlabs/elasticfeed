package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resource"
	"github.com/feedlabs/api/helper"
	"github.com/feedlabs/feedify/service"
)

type DefaultController struct {
	feedify.Controller
}

var (
	Admin *resource.Admin
)

func (this *DefaultController) Get() {
	this.Data["json"] = map[string]string{"succes": "ok"}
	this.ServeJson()
}

func SetGlobalResponseHeader() {
	var FilterUser = func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	var AuthUser = func(ctx *context.Context) {
		Admin = helper.Auth(ctx)
	}
	beego.InsertFilter("/*", beego.BeforeRouter, AuthUser)
}

func GetMyOrgId() string {
	if Admin != nil && Admin.Org != nil {
		return Admin.Org.Id
	}
	return "0"
}

func AdminChannelID() string {
	return helper.AdminChannelID(Admin)
}

func init() {
	SetGlobalResponseHeader()
	graph, _ := service.NewGraph()
	graph.Storage.Connect()
}
