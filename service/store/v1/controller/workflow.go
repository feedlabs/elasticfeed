package controller

import (
	"encoding/json"

	"github.com/feedlabs/elasticfeed/resource"
	template "github.com/feedlabs/elasticfeed/service/store/v1/template/workflow"
)

type WorkflowController struct {
	DefaultController
}

/**
 * @api {get} application/:applicationId/feed/:feedId/workflow Get List (Feed)
 * @apiVersion 1.0.0
 * @apiName GetWorkflowListFeed
 * @apiGroup Workflow
 * @apiDescription This will return a list of all entries per feed you have created.
 *
 * @apiUse WorkflowGetListRequest
 * @apiUse WorkflowGetListResponse
 */
func (this *WorkflowController) GetList() {
	formatter, err := template.RequestGetList(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feed, err := resource.GetFeed(feedId, appId, this.GetAdminOrgId())

	obs, err := feed.GetWorkflowList()
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponseGetList(obs, formatter))
	}
}

/**
 * @api {get} application/:applicationId/feed/:feedId/workflow/:workflowId Get
 * @apiVersion 1.0.0
 * @apiName GetWorkflow
 * @apiGroup Workflow
 * @apiDescription This will return a specific workflow.
 *
 * @apiUse WorkflowGetRequest
 * @apiUse WorkflowGetResponse
 */
func (this *WorkflowController) Get() {
	formatter, err := template.RequestGet(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feedWorkflowId := this.Ctx.Input.Params[":feedWorkflowId"]

	ob, err := resource.GetWorkflow(feedWorkflowId, feedId, appId, this.GetAdminOrgId())
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponseGet(ob, formatter))
	}
}

/**
 *
 * @api {post} application/:applicationId/feed/:feedId/workflow/ Add (Feed)
 * @apiVersion 1.0.0
 * @apiName PostWorkflowFeedAdd
 * @apiGroup Workflow
 * @apiDescription Add a workflow to the feed which is already store in the system.
 *
 * @apiUse WorkflowAddToFeedRequest
 * @apiUse WorkflowAddToFeedResponse
 */
func (this *WorkflowController) Post() {
	formatter, err := template.RequestPost(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]

	var ob resource.Workflow
	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	app, err := resource.GetApplication(appId, this.GetAdminOrgId())
	feed, err := app.GetFeed(feedId)

	_, err = feed.AddWorkflow(ob)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponsePost(&ob, formatter))
	}
}

/**
 * @api {put} application/:applicationId/workflow/:workflowId Update (Global)
 * @apiVersion 1.0.0
 * @apiName PutWorkflow
 * @apiGroup Workflow
 * @apiDescription Update a specific workflow.
 *
 * @apiUse WorkflowPutRequest
 * @apiUse WorkflowPutResponse
 */
func (this *WorkflowController) Put() {
	formatter, err := template.RequestPut(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feedWorkflowId := this.Ctx.Input.Params[":feedWorkflowId"]

	var ob *resource.Workflow

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	_ob, err := resource.UpdateWorkflow(feedWorkflowId, feedId, appId, this.GetAdminOrgId(), ob.Data)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponsePut(_ob, formatter))
	}
}

/**
 * @api {delete} application/:applicationId/workflow/:workflowId Delete (Global)
 * @apiVersion 1.0.0
 * @apiName DeleteWorkflow
 * @apiGroup Workflow
 * @apiDescription Delete a specific workflow. (will also remove the workflow from all feeds)
 *
 * @apiUse WorkflowDeleteRequest
 * @apiUse WorkflowDeleteResponse
 */
func (this *WorkflowController) Delete() {
	formatter, err := template.RequestDelete(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feedWorkflowId := this.Ctx.Input.Params[":feedWorkflowId"]

	err = resource.DeleteWorkflow(feedWorkflowId, feedId, appId, this.GetAdminOrgId())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	} else {
		this.ServeJson(template.ResponseDelete("Org has been deleted", formatter))
	}
}
