package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/public/v1/controller"
	"github.com/feedlabs/api/public/v1/controller/org"
)

func init() {
	feedify.Router("/v1/org", &controller.OrgController{}, "get:GetList;post:Post")
	feedify.Router("/v1/org/:orgId:string", &controller.OrgController{}, "get:Get;delete:Delete;put:Put")

	feedify.Router("/v1/org/:orgId:string/token", &org.TokenController{}, "get:GetList;post:Post")
	feedify.Router("/v1/org/:orgId:string/token/:tokenId:string", &org.TokenController{}, "get:Get;delete:Delete")
}
