package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/public/v1/controller"
	"github.com/feedlabs/api/public/v1/controller/admin"
)

func init() {
	feedify.Router("/v1/admin", &controller.AdminController{}, "get:GetList;post:Post")
	feedify.Router("/v1/admin/:adminId:string", &controller.AdminController{}, "get:Get;delete:Delete;put:Put")

	feedify.Router("/v1/admin/:adminId:string/token", &admin.TokenController{}, "get:GetList;post:Post")
	feedify.Router("/v1/admin/:adminId:string/token/:tokenId:string", &admin.TokenController{}, "get:Get;delete:Delete")
}
