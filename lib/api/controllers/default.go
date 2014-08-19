package controllers

import (
	"strings"
	"github.com/feedlabs/feedify/lib/feedify/db/adapter"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type ResponseInfo struct {
}

var mongo *adapter.Mongo
var neo4j *adapter.Neo4j
var memcache *adapter.Memcache
var cayley *adapter.Cayley

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
	mongo = adapter.NewMongo()
	mongo.Connect()

	neo4j = adapter.NewNeo4j()
	neo4j.Connect()

	memcache = adapter.NewMemcache()
	memcache.Connect()

	cayley = adapter.NewCayley()
	cayley.Connect()
}
