package routers

import (
	"cfp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/storage", &controllers.StorageController{}, "get:Get;post:Post;options:Options")
	beego.Router("/v1/storage/:id:int", &controllers.StorageController{}, "get:Get;delete:Delete;put:Put")
}
