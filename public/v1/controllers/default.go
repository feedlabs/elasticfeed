package controllers

import (
	"fmt"

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

func init() {
	graph, _ := service.NewGraph()
	graph.Storage.Connect()
	graph.Storage.Query(`
				START n=node(*)
				MATCH (n)-[r:outranks]->(m)
				WHERE n.shirt = {color}
				RETURN n.name, type(r), m.name
			`)

	memcache := service.NewCache()
	memcache.Connect()
	memcache.Set("feed", "hello from memcache feed")
	memcache.Set("api", "hello from memcache api")
	fmt.Println(memcache.GetMulti([]string{"feed", "api"}))
}
