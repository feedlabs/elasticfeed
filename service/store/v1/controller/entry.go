package controller

import (
	"encoding/json"

	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/service/store/v1/template/entry"
)

type EntryController struct {
	DefaultController
}

/**
 * @api {get} application/:applicationId/feed/:feedId/entry Get List (Feed)
 * @apiVersion 1.0.0
 * @apiName GetEntryListFeed
 * @apiGroup Entry
 * @apiDescription This will return a list of all entries per feed you have created.
 *
 * @apiUse EntryGetListByFeedRequest
 * @apiUse EntryGetListByFeedResponse
 */
func (this *EntryController) GetList() {
	entry.RequestGetListByFeed(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feed, err := resource.GetFeed(feedId, appId, this.GetAdminOrgId())
	obs, err := feed.GetEntryList()

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	entry.ResponseGetListByFeed()
	this.Controller.ServeJson()
}

/**
 * @api {get} application/:applicationId/feed/:feedId/entry/:entryId Get (Global)
 * @apiVersion 1.0.0
 * @apiName GetEntry
 * @apiGroup Entry
 * @apiDescription This will return a specific entry.
 *
 * @apiUse EntryGetRequest
 * @apiUse EntryGetResponse
 */
func (this *EntryController) Get() {
	// two different usages both
	// 1: get entry from specific feed (includes feedId)
	// 2: get (global) entry from application

	entry.RequestGet(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	ob, err := resource.GetEntry(feedEntryId, feedId, appId, this.GetAdminOrgId())
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = ob
	}

	entry.ResponseGet()
	this.Controller.ServeJson()
}

/**
 *
 * @api {post} application/:applicationId/feed/:feedId/entry/ Add (Feed)
 * @apiVersion 1.0.0
 * @apiName PostEntryFeedAdd
 * @apiGroup Entry
 * @apiDescription Add a entry by entry Id to the feed which is already store in the system.
 *
 * @apiUse EntryAddToFeedRequest
 * @apiUse EntryAddToFeedResponse
 */
func (this *EntryController) Post() {
	// two different usages
	// 1: post new data and create entry directly to feed
	// 2: post just a entryId which will be added to feed

	entry.RequestPost(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]

	var ob resource.Entry
	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	app, err := resource.GetApplication(appId, this.GetAdminOrgId())
	feed, err := app.GetFeed(feedId)
	entryId, err := feed.AddEntry(ob)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": entryId}
	}

	entry.ResponsePost()
	this.Controller.ServeJson()
}

/**
 * @api {put} application/:applicationId/entry/:entryId Update (Global)
 * @apiVersion 1.0.0
 * @apiName PutEntry
 * @apiGroup Entry
 * @apiDescription Update a specific entry.
 *
 * @apiUse EntryPutRequest
 * @apiUse EntryPutResponse
 */
func (this *EntryController) Put() {
	entry.RequestPut(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	var ob resource.Entry

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	err := resource.UpdateEntry(feedEntryId, feedId, appId, this.GetAdminOrgId(), ob.Data)
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "update success", "status": "ok"}
	}

	entry.ResponsePut()
	this.Controller.ServeJson()
}

/**
 * @api {delete} application/:applicationId/feed/:feedId/entry/:entryId Delete (Global)
 * @apiVersion 1.0.0
 * @apiName DeleteEntry
 * @apiGroup Entry
 * @apiDescription Delete a specific entry. (will also remove the entry from all feeds)
 *
 * @apiUse EntryDeleteRequest
 * @apiUse EntryDeleteResponse
 */
func (this *EntryController) Delete() {
	entry.RequestDelete(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	err := resource.DeleteEntry(feedEntryId, feedId, appId, this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	}

	entry.ResponseDelete()
	this.Controller.ServeJson()
}

