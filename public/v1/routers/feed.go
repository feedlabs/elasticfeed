package routers

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/public/v1/controllers"
)

func init() {
	feedify.Router("/v1/feed", &controllers.FeedController{}, "get:Get;post:Post")
	feedify.Router("/v1/feed/:feedId:int", &controllers.FeedController{}, "get:Get;delete:Delete;put:Put")

	feedify.Router("/v1/feed/:feedId:int/entry", &controllers.FeedEntryController{}, "get:Get;post:Post")
	feedify.Router("/v1/feed/:feedId:int/entry/:feedEntryId:int", &controllers.FeedEntryController{}, "get:Get;delete:Delete;put:Put")
}
