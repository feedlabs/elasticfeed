package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/system/v1/controller"
)

func InitRouters() {
	feedify.Router("/v1/system/status", &controller.StatusController{}, "get:Get")
}
