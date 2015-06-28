package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/surrogate/v1/controller"
)

func InitTrainRouters() {
	feedify.Router("/v1/surrogate/train", &controller.DefaultController{}, "get:Get")
}
