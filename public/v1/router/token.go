package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/public/v1/controller"
)

func init() {
	feedify.Router("/v1/admin/:adminId:string/token", &controller.TokenController{}, "get:GetList;post:Post")
	feedify.Router("/v1/admin/:adminId:string/token/:tokenId:string", &controller.TokenController{}, "get:Get;delete:Delete")
}
