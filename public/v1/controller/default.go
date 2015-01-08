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

func (this *DefaultController) Get() {
	this.Data["json"] = map[string]string{"succes": "ok"}
	this.ServeJson()
}

func SetGlobalResponseHeader() {
	var FilterUser = func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
}

func GetMyOrgId() string {
	return resources.GetClientOrgId()
}

func GetSecret() string {
	return resources.GetApiSecret()
}

func AuthenticateHTTPRequest() {
	// should handle auth
	// get token length
	// 32 -> client token -> orgId
	// 128 -> global api token
	// IP whitelisting
}

func GenerateChannelID() {
	// should generate proper ID
}

func init() {
	SetGlobalResponseHeader()
	graph, _ := service.NewGraph()
	graph.Storage.Connect()
}
