package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/system/controller"
)

func InitOrgRouters() {
	feedify.Router("/v1/plugin", &controller.PluginController{}, "get:GetList;post:Post")
	feedify.Router("/v1/plugin/:pluginId:string", &controller.PluginController{}, "get:Get;delete:Delete;put:Put")
}
