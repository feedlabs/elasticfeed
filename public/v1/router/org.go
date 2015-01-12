package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/public/v1/controller"
)

func init() {
	feedify.Router("/v1/org", &controller.OrgController{}, "get:GetList;post:Post")
	feedify.Router("/v1/org/:orgId:string", &controller.OrgController{}, "get:Get;delete:Delete;put:Put")
}
