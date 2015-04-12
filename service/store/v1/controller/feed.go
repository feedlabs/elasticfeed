package controller

import (
	"encoding/json"

	"github.com/feedlabs/elasticfeed/resource"

	"github.com/feedlabs/elasticfeed/service/store/v1/template"
	"github.com/feedlabs/elasticfeed/service/store/v1/template/feed"
)

type FeedController struct {
	DefaultController
}

/**
 * @api {get} application/:applicationId/feed Get List
 * @apiVersion 1.0.0
 * @apiName GetFeedList
 * @apiGroup Feed
 * @apiDescription This will return a list of all feeds per applications you have created.
 *
 * @apiUse FeedGetListRequest
 * @apiUse FeedGetListResponse
 */
func (this *FeedController) GetList() {
	feed.RequestGetList(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	app, err := resource.GetApplication(appId, this.GetAdminOrgId())
	obs, err := app.GetFeedList()

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	feed.ResponseGetList()
	this.Controller.ServeJson()
}

/**
 * @api {get} application/:applicationId/feed/:feedId Get
 * @apiVersion 1.0.0
 * @apiName GetFeed
 * @apiGroup Feed
 * @apiDescription This will return a specific feed.
 *
 * @apiUse FeedGetRequest
 * @apiUse FeedGetResponse
 */
func (this *FeedController) Get() {
	feed.RequestGet(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	ob, err := resource.GetFeed(feedId, appId, this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = ob
	}

	feed.ResponseGet()
	this.Controller.ServeJson()
}

/**
 * @api {post} application/:applicationId/feed Create
 * @apiVersion 1.0.0
 * @apiName PostFeed
 * @apiGroup Feed
 * @apiDescription Create a feed.
 *
 * @apiUse FeedPostRequest
 * @apiUse FeedPostResponse
 */
func (this *FeedController) Post() {
	feed.RequestPost(this.GetInput())

	var ob resource.Feed

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	appId := this.Ctx.Input.Params[":applicationId"]
	app, err := resource.GetApplication(appId, this.GetAdminOrgId())
	feedId, err := app.AddFeed(ob)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": feedId}
	}

	feed.ResponsePost()
	this.Controller.ServeJson()
}

/**
 * @api {put} application/:applicationId/feed/:feedId Update
 * @apiVersion 1.0.0
 * @apiName PutFeed
 * @apiGroup Feed
 * @apiDescription Update a specific feed.
 *
 * @apiUse FeedPostRequest
 * @apiUse FeedPostResponse
 */
func (this *FeedController) Put() {
	feed.RequestPut(this.GetInput())

	feedId := this.Ctx.Input.Params[":feedId"]
	var ob resource.Feed

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	err := resource.UpdateFeed(feedId, ob.Data)
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "update success", "status": "ok"}
	}

	feed.ResponsePut()
	this.Controller.ServeJson()
}

/**
 * @api {delete} application/:applicationId/feed/:feedId Delete
 * @apiVersion 1.0.0
 * @apiName DeleteFeed
 * @apiGroup Feed
 * @apiDescription Delete a specific feed.
 *
 * @apiUse FeedDeleteRequest
 * @apiUse FeedDeleteResponse
 */
func (this *FeedController) Delete() {
	feed.RequestDelete(this.GetInput())

	feedId := this.Ctx.Input.Params[":feedId"]
	err := resource.DeleteFeed(feedId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	}

	feed.ResponseDelete()
	this.Controller.ServeJson()
}

func (this *FeedController) ActionReload() {
	feedId := this.Ctx.Input.Params[":feedId"]

	resource.ActionReloadFeed(feedId)

	this.ServeJson(template.Success("Feed reloaded."))
}

func (this *FeedController) ActionEmpty() {
	feedId := this.Ctx.Input.Params[":feedId"]

	resource.ActionEmptyFeed(feedId)

	this.ServeJson(template.Success("Feed empty done."))
}
