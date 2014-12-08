package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/feedify/service"
)

type ResponseInfo struct {
}

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

func init() {
	SetGlobalResponseHeader()
	graph, _ := service.NewGraph()
	graph.Storage.Connect()
}
