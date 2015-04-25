package router

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/elasticfeed/service/store/v1/controller"
)

func InitEntryWorkflows() {
	feedify.Router("/v1/application/:applicationId:string/feed/:feedId:int/workflow", &controller.WorkflowController{}, "get:GetList;post:Post")
	feedify.Router("/v1/application/:applicationId:string/feed/:feedId:int/workflow/:feedWorkflowId:int", &controller.WorkflowController{}, "get:Get;delete:Delete;put:Put")
}
