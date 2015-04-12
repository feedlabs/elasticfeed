package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/system/controller"
)

func InitRouters() {
	feedify.Router("/system/status", &controller.StatusController{}, "get:Get")
}
