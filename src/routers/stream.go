package routers

import (
	"github.com/feedlabs/feedify/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/stream", &controllers.StreamController{}, "get:Get;post:Post")
	beego.Router("/v1/stream/:id:int", &controllers.StreamController{}, "get:Get;delete:Delete;put:Put")
}
