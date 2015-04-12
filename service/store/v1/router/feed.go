package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/store/v1/controller"
)

func InitFeedRouters() {
	feedify.Router("/v1/application/:applicationId:string/feed", &controller.FeedController{}, "get:GetList;post:Post")
	feedify.Router("/v1/application/:applicationId:string/feed/:feedId:int", &controller.FeedController{}, "get:Get;delete:Delete;put:Put")

	feedify.Router("/v1/application/:applicationId:string/feed/:feedId:int/reload", &controller.FeedController{}, "get:ActionReload")
	feedify.Router("/v1/application/:applicationId:string/feed/:feedId:int/empty", &controller.FeedController{}, "get:ActionEmpty")
}
