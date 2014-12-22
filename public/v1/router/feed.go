package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/public/v1/controller"
)

func init() {
	feedify.Router("/v1/application/:applicationId:string/feed", &controller.FeedController{}, "get:GetList;post:Post")
	feedify.Router("/v1/application/:applicationId:string/feed/:feedId:int", &controller.FeedController{}, "get:Get;delete:Delete;put:Put")
}
