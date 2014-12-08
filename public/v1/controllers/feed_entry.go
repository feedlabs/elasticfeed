package controllers

import (
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
)

type FeedEntryController struct {
	feedify.Controller
}

func (this *FeedEntryController) Post() {
	var ob resources.FeedEntry

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	feedId := this.Ctx.Input.Params[":feedId"]
	id, err := resources.AddFeedEntry(ob, feedId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": id}
	}

	this.ServeJson()
}

func (this *FeedEntryController) Get() {
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	if feedId != "" && feedEntryId != "" {
		ob, err := resources.GetFeedEntry(feedEntryId, feedId)
		if err != nil {
			this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs, err := resources.GetFeedEntryList(feedId)
		this.Data["json"] = obs
		if err != nil {
			this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
		} else {
			this.Data["json"] = obs
		}
	}

	this.ServeJson()
}

func (this *FeedEntryController) Put() {
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	var ob resources.FeedEntry

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	err := resources.UpdateFeedEntry(feedEntryId, feedId, ob.Data)
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "update success", "status": "ok"}
	}
	this.ServeJson()
}

func (this *FeedEntryController) Delete() {
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	err := resources.DeleteFeedEntry(feedEntryId, feedId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	}

	this.ServeJson()
}
