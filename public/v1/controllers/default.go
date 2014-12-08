package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/feedify/service"
)

const VALID_REQUEST = 200
const ENTITY_CREATED = 201
const ENTITY_NOEXIST = 404
const ENTITY_CONFLICT = 409
const INVALID_REQUEST = 400
const ACCESS_UNAUTHORIZED = 401
const ACCESS_FORBIDDEN = 403
const NOALLOWED_REQUEST = 405
const TOOMANY_REQUEST = 429
const SERVER_ERROR = 500

type ResponseInfo struct {
}

type DefaultController struct {
	feedify.Controller
}

type DefaultGETRequest struct {
}

type DefaultPOSTRequest struct {

}

type DefaultPUTRequest struct {

}

type DefaultDELETERequest struct {

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

func ValidateHTTPHeaderInput(template_controller int, request_input int) {
	// ctx.Input.Header()
}

func ValidateHTTPBodyInput() {
	// should be done for JSON
	// http://play.golang.org/p/_bKAQ3dQlu
	// http://blog.golang.org/laws-of-reflection
	// http://stackoverflow.com/questions/18926303/iterate-through-a-struct-in-go
}

func ValidateHTTPParamsInput() {
	// ctx.Input.Params
}

func AuthenticateHTTPRequest() {

}

func SetHTTPHeaderOutput(code int) {
	// set HTTP response CODE
	// ctx.Output.SetStatus(200)
}

func ValidateHTTPOutput() {

}

func PrettyPrint() {
	// ctx.Input.Params("pretty") == true ?
}

func LowerJSONKeys() {

}

func GenerateChannelID() {

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
	query := graph.Storage.Query(`
				START n=node(*)
				WHERE n:feed
				RETURN n
			`)

	fmt.Println(query.Result)

	memcache := service.NewCache()
	memcache.Connect()
	memcache.Set("feed", "hello from memcache feed")
	memcache.Set("api", "hello from memcache api")
	fmt.Println(memcache.GetMulti([]string{"feed", "api"}))
}
