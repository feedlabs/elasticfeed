package controllers

import (
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
	"github.com/feedlabs/api/public/v1/template/feed"
)

type FeedController struct {
	feedify.Controller
}

func (this *FeedController) Post() {
	var ob resources.Feed

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	feedid, err := resources.AddFeed(ob)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": feedid}
	}

	this.ServeJson()
}

/**
 * @api {get} application/:applicationId/feed Get List
 * @apiVersion 1.0.0
 * @apiName GetFeedList
 * @apiGroup Feed
 *
 * @apiDescription This will return a list of all feeds per applications you have created.
 *
 * @apiUse FeedGetListRequest
 * @apiUse FeedGetListResponse
 */
func (this *FeedController) GetList() {
	feed.RequestGetList(this.GetInput())

	obs := resources.GetFeedList()
	this.Data["json"] = obs

	feed.ResponseGetList()
	this.ServeJson()
}

/**
 * @api {get} application/:applicationId/feed/:feedId Get
 * @apiVersion 1.0.0
 * @apiName GetFeed
 * @apiGroup Feed
 *
 * @apiDescription This will return a specific feed.
 *
 * @apiUse FeedGetRequest
 * @apiUse FeedGetResponse
 */
func (this *FeedController) Get() {
	feed.RequestGet(this.GetInput())

	feedId := this.Ctx.Input.Params[":feedId"]
	ob, err := resources.GetFeed(feedId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = ob
	}

	feed.ResponseGet()
	this.ServeJson()
}

func (this *FeedController) Put() {
	feedId := this.Ctx.Input.Params[":feedId"]
	var ob resources.Feed

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	err := resources.UpdateFeed(feedId, ob.Data)
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "update success", "status": "ok"}
	}
	this.ServeJson()
}

func (this *FeedController) Delete() {
	feedId := this.Ctx.Input.Params[":feedId"]
	err := resources.DeleteFeed(feedId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	}

	this.ServeJson()
}
