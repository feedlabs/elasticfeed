package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/system/controller"
)

func init() {
	feedify.Router("/system/status", &controller.StatusController{}, "get:Get")
}
