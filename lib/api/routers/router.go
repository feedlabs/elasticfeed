package routers

import (
	"github.com/astaxie/beego"
	"github.com/feedlabs/feedify/lib/api/controllers"
)

func init() {
	beego.Router("/", &controllers.DefaultController{}, "get:Get")
	beego.Router("/v1", &controllers.DefaultController{}, "get:Get")
}
