package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/common"
	"github.com/feedlabs/elasticfeed/service/store/v1/template"
)

type DefaultController struct {
	feedify.Controller
}

func (this *DefaultController) Get() {
	this.Data["json"] = map[string]string{"succes": "ok"}
	this.Controller.ServeJson()
}

func (this *DefaultController) GetAdmin() *resource.Admin {
	if this.GetCtx().Input.Data["admin"] != nil {
		return this.GetCtx().Input.Data["admin"].(*resource.Admin)
	}
	return nil
}

func (this *DefaultController) GetAdminOrgId() string {
	admin := this.GetAdmin()
	if admin != nil {
		return admin.Org.Id
	}
	return "0"
}

func (this *DefaultController) ServeJson(data interface{}, status int) {
	this.Data["json"] = data
	this.SetResponseStatusCode(status)
	this.Controller.ServeJson()
}

func (this *DefaultController) SetResponseStatusCode(code int) {
	this.Controller.Ctx.Output.SetStatus(code)
}

func SetGlobalResponseHeader() {
	var FilterUser = func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
}

func SetAuthenticationFilter() {
	var AuthUser = func(ctx *context.Context) {
		ctx.Input.Data["admin"] = common.Auth(ctx)
	}
	beego.InsertFilter("/v1/*", beego.BeforeRouter, AuthUser)
}

func NoRoutes() {
	var router = func(ctx *context.Context) {
		if ctx.Output.Status == 0 {
			ctx.Output.SetStatus(template.HTTP_CODE_INVALID_REQUEST)
		}
	}
	beego.InsertFilter("/*", beego.AfterStatic, router)
}

func InitService() {
	SetAuthenticationFilter()
	SetGlobalResponseHeader()
	//	NoRoutes()
}
