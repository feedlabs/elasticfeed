package routers

import (
	"github.com/feedlabs/feedify/lib/api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/feed", &controllers.FeedController{}, "get:Get;post:Post")
	beego.Router("/v1/feed/:id:int", &controllers.FeedController{}, "get:Get;delete:Delete;put:Put")
}
