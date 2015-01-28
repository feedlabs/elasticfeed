package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/api/v1/controller"
)

func init() {
	feedify.Router("/service/api/v1/org/:orgId:string/token", &controller.TokenController{}, "get:GetOrgList;post:PostOrg")
	feedify.Router("/service/api/v1/org/:orgId:string/token/:tokenId:string", &controller.TokenController{}, "get:GetOrg;delete:DeleteOrg")

	feedify.Router("/service/api/v1/admin/:adminId:string/token", &controller.TokenController{}, "get:GetAdminList;post:PostAdmin")
	feedify.Router("/service/api/v1/admin/:adminId:string/token/:tokenId:string", &controller.TokenController{}, "get:GetAdmin;delete:DeleteAdmin")
}
