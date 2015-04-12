package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/store/v1/controller"
)

func InitApplicationRouters() {
	feedify.Router("/v1/application", &controller.ApplicationController{}, "get:GetList;post:Post")
	feedify.Router("/v1/application/:applicationId:string", &controller.ApplicationController{}, "get:Get;delete:Delete;put:Put")
}
