package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/api/v1/controller"
)

func init() {
	feedify.Router("/", &controller.DefaultController{}, "get:Get")
	feedify.Router("/v1", &controller.DefaultController{}, "get:Get")
}
