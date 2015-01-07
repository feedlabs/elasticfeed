package controller

import (
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
	"github.com/feedlabs/api/public/v1/template/entry"
)

type EntryController struct {
	feedify.Controller
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
func (this *EntryController) GetListByFeed() {
	entry.RequestGetListByFeed(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feed, err := resources.GetFeed(feedId, appId)
	obs, err := feed.GetEntryList()

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	entry.ResponseGetListByFeed()
	this.ServeJson()
}

/**
 * @api {get} application/:applicationId/entry/:entryId Get
 * @apiVersion 1.0.0
 * @apiName GetEntry
 * @apiGroup Entry
 * @apiDescription This will return a specific entry.
 *
 * @apiUse EntryGetRequest
 * @apiUse EntryGetResponse
 */
/**
 * @api {get} application/:applicationId/feed/:feedId/entry/:entryId Get (Feed)
 * @apiVersion 1.0.0
 * @apiName GetEntryFeed
 * @apiGroup Entry
 * @apiDescription This will return a specific entry.
 *
 * @apiUse EntryGetByFeedRequest
 * @apiUse EntryGetByFeedResponse
 */
func (this *EntryController) Get() {
	// two different usages both
	// 1: get entry from specific feed (includes feedId)
	// 2: get (global) entry from application

	entry.RequestGet(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	ob, err := resources.GetEntry(feedEntryId, feedId, appId)
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = ob
	}

	entry.ResponseGet()
	this.ServeJson()
}

/**
 * @api {post} application/:applicationId/entry Create
 * @apiVersion 1.0.0
 * @apiName PostEntry
 * @apiGroup Entry
 * @apiDescription Create a entry on the global feed. This could be used to store a element in the cloud system and re-use it later.
 *
 * @apiUse EntryPostRequest
 * @apiUse EntryPostResponse
 */
func (this *EntryController) Post() {
	// global entry; should be added to APP no to the FEED
	entry.RequestPost(this.GetInput())
	entry.ResponsePost()
	this.ServeJson()
}

/**
 * @api {post} application/:applicationId/feed/:feedId/entry Create (Feed)
 * @apiVersion 1.0.0
 * @apiName PostEntryFeed
 * @apiGroup Entry
 * @apiDescription Create a entry in the global feed and link it automatically to a feed.
 *
 * @apiUse EntryPostToFeedRequest
 * @apiUse EntryPostToFeedResponse
 */
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
func (this *EntryController) PostToFeed() {
	// two different usages
	// 1: post new data and create entry directly to feed
	// 2: post just a entryId which will be added to feed

	entry.RequestPost(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	feedId := this.Ctx.Input.Params[":feedId"]

	var ob resources.Entry
	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	app, err := resources.GetApplication(appId)
	feed, err := app.GetFeed(feedId)
	entryId, err := feed.AddEntry(ob)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": entryId}
	}

	entry.ResponsePost()
	this.ServeJson()
}

/**
 * @api {put} application/:applicationId/entry/:entryId Update
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

	var ob resources.Entry

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	err := resources.UpdateEntry(feedEntryId, feedId, appId, ob.Data)
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "update success", "status": "ok"}
	}

	entry.ResponsePut()
	this.ServeJson()
}

/**
 * @api {delete} application/:applicationId/entry/:entryId Delete
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

	err := resources.DeleteEntry(feedEntryId, feedId, appId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	}

	entry.ResponseDelete()
	this.ServeJson()
}

/**
 * @api {delete} application/:applicationId/feed/:feedId/entry/:entryId Remove (Feed)
 * @apiVersion 1.0.0
 * @apiName RemoveEntry
 * @apiGroup Entry
 * @apiDescription Removes a specific entry from a feed.
 *
 * @apiUse EntryRemoveRequest
 * @apiUse EntryRemoveResponse
 */
func (this *EntryController) Remove() {
	entry.RequestRemove(this.GetInput())
	entry.ResponseRemove()
	this.ServeJson()
}
