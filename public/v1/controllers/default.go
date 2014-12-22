package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/feedlabs/feedify"
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

func AuthenticateHTTPRequest() {
	// should handle auth
}

func GenerateChannelID() {
	// should generate proper ID
}

func GenerateFeedID() {
	// should contain channelID
}

func GenerateClientID() {
	// Should be as base for feedID and feedChannelID
	// clientID should allow to generate single channel (websocket connection) for multiple feed-pages
	// if used public/private multiple feed-pages in the same time there should be up to 2 websocket connections
}

func init() {
	SetGlobalResponseHeader()
	graph, _ := service.NewGraph()
	graph.Storage.Connect()
}
