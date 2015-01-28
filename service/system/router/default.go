package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/system/controller"
)

func init() {
	feedify.Router("/status", &controller.SystemController{}, "get:Get")
}
