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
	UserOrg *resource.Org
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
		UserOrg = helper.Auth(ctx)
	}
	beego.InsertFilter("/*", beego.BeforeRouter, AuthUser)
}

func GetMyOrgId() string {
	return UserOrg.Id
}

func GetSecret() string {
	return helper.GetApiSecret()
}

func GenerateChannelID() string {
	return helper.GenerateChannelID()
}

func init() {
	SetGlobalResponseHeader()
	graph, _ := service.NewGraph()
	graph.Storage.Connect()
}
