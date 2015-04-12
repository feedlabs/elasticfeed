package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/store/v1/controller"
)

func InitOrgRouters() {
	feedify.Router("/v1/org", &controller.OrgController{}, "get:GetList;post:Post")
	feedify.Router("/v1/org/:orgId:string", &controller.OrgController{}, "get:Get;delete:Delete;put:Put")
}
