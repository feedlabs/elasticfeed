package controller

import (
	"github.com/feedlabs/feedify"
)

type StatusController struct {
	feedify.Controller
}

func (this *StatusController) Get() {
	this.Data["json"] = map[string]interface{}{
		"enabled": true,
	}

	this.Controller.ServeJson()
}
