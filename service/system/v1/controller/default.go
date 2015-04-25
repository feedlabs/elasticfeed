package controller

import (
	"github.com/feedlabs/feedify"
)

type DefaultController struct {
	feedify.Controller
}

func (this *DefaultController) Get() {
	this.Data["json"] = map[string]string{"succes": "ok"}
	this.Controller.ServeJson()
}

func (this *DefaultController) ServeJson(data interface{}, status int) {
	this.Data["json"] = data
	this.SetResponseStatusCode(status)
	this.Controller.ServeJson()
}

func (this *DefaultController) SetResponseStatusCode(code int) {
	this.Controller.Ctx.Output.SetStatus(code)
}
