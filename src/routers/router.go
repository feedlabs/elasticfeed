package routers

import (
	"cfp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.DefaultController{}, "get:Get")
	beego.Router("/v1", &controllers.DefaultController{}, "get:Get")
}
