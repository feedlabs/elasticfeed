package routers

import (
	"github.com/astaxie/beego"
	"github.com/feedlabs/api/public/v1/controllers"
)

func init() {
	beego.Router("/v1/feed", &controllers.FeedController{}, "get:Get;post:Post")
	beego.Router("/v1/feed/:feedId:int", &controllers.FeedController{}, "get:Get;delete:Delete;put:Put")

	beego.Router("/v1/feed/:feedId:int/entry", &controllers.FeedEntryController{}, "get:Get;post:Post")
	beego.Router("/v1/feed/:feedId:int/entry/:feedEntryId:int", &controllers.FeedEntryController{}, "get:Get;delete:Delete;put:Put")
}
