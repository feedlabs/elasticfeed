package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/api/v1/controller"
)

func init() {
	feedify.Router("/v1/admin", &controller.AdminController{}, "get:GetList;post:Post")
	feedify.Router("/v1/admin/:adminId:string", &controller.AdminController{}, "get:Get;delete:Delete;put:Put")
}
