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

func (this *DefaultController) GetAdmin() *resource.Admin {
	if this.GetCtx().Input.Data["admin"] != nil {
		return this.GetCtx().Input.Data["admin"].(*resource.Admin)
	}
	return nil
}

func (this *DefaultController) GetAdminOrgId() string {
	admin := this.GetAdmin()
	if admin != nil {
		return admin.Org.Id
	}
	return "0"
}

func SetGlobalResponseHeader() {
	var FilterUser = func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	var AuthUser = func(ctx *context.Context) {
		ctx.Input.Data["admin"] = helper.Auth(ctx)
	}
	beego.InsertFilter("/*", beego.BeforeRouter, AuthUser)
}

func AdminChannelID() string {
	return helper.AdminChannelID(Admin)
}

func init() {
	SetGlobalResponseHeader()
	graph, _ := service.NewGraph()
	graph.Storage.Connect()
}
