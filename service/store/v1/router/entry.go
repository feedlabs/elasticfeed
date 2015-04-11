package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/store/v1/controller"
)

func InitEntryRouters() {
	feedify.Router("/v1/application/:applicationId:string/entry", &controller.EntryController{}, "post:Post")
	feedify.Router("/v1/application/:applicationId:string/entry/:feedEntryId:int", &controller.EntryController{}, "get:Get;delete:Delete;put:Put")

	feedify.Router("/v1/application/:applicationId:string/feed/:feedId:int/entry", &controller.EntryController{}, "get:GetListByFeed;post:PostToFeed")
	feedify.Router("/v1/application/:applicationId:string/feed/:feedId:int/entry/:feedEntryId:int", &controller.EntryController{}, "get:Get;delete:Remove;put:Put")
}
