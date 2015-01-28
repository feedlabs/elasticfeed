package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/api/v1/controller"
)

func init() {
	feedify.Router("/service/api/v1/application", &controller.ApplicationController{}, "get:GetList;post:Post")
	feedify.Router("/service/api/v1/application/:applicationId:string", &controller.ApplicationController{}, "get:Get;delete:Delete;put:Put")
}
