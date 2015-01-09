package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
	"github.com/feedlabs/feedify/service"
)

type DefaultController struct {
	feedify.Controller
}

var (
	UserOrg *resources.Org
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
		UserOrg = resources.Auth(ctx)
	}
	beego.InsertFilter("/*", beego.BeforeRouter, AuthUser)
}

func GetMyOrgId() string {
	return UserOrg.Id
}

func GetSecret() string {
	return resources.GetApiSecret()
}

func GenerateChannelID() {
	// should generate proper ID
}

func init() {
	SetGlobalResponseHeader()
	graph, _ := service.NewGraph()
	graph.Storage.Connect()
}
