package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/predict/v1/controller"
)

func InitPredictRouters() {
	feedify.Router("/v1/predict/predict", &controller.DefaultController{}, "get:Get")
}
