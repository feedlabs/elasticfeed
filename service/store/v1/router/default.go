package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/store/v1/controller"
)

func InitDefaultRouters() {
	feedify.Router("/", &controller.DefaultController{}, "get:Get")
	feedify.Router("/v1", &controller.DefaultController{}, "get:Get")
}

func InitRouters() {
	InitDefaultRouters()
	InitAdminRouters()
	InitApplicationRouters()
	InitEntryRouters()
	InitEntryWorkflows()
	InitFeedRouters()
	InitOrgRouters()
	InitTokenRouters()
}
