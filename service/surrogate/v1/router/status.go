package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/surrogate/v1/controller"
)

func InitStatusRouters() {
	feedify.Router("/v1/surrogate/status", &controller.StatusController{}, "get:Get")
}
