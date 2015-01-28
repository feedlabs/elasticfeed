package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/db/v1/controller"
)

func init() {
	feedify.Router("/", &controller.DefaultController{}, "get:Get")
	feedify.Router("/db/", &controller.DefaultController{}, "get:Get")
	feedify.Router("/db/v1", &controller.DefaultController{}, "get:Get")
}
