package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/surrogate/v1/controller"
)

func InitQueryRouters() {
	feedify.Router("/v1/surrogate/query", &controller.DefaultController{}, "get:Get")
}
