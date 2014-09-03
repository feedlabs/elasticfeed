package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/feedlabs/feedify/lib/feedify/entity"
)

type FeedEntryController struct {
	beego.Controller
}

func (this *FeedEntryController) Post() {
	var ob entity.FeedEntry

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	feedId := this.Ctx.Input.Params[":feedId"]

	id := entity.AddFeedEntry(ob, feedId)
	this.Data["json"] = map[string]string{"id": id}
	this.ServeJson()
}

func (this *FeedEntryController) Get() {
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	if feedId != "" && feedEntryId != "" {
		ob, err := entity.GetFeedEntry(feedEntryId, feedId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := entity.GetFeedEntryList(feedId)
		this.Data["json"] = obs
	}

	this.ServeJson()
}

func (this *FeedEntryController) Put() {
	feedId := this.Ctx.Input.Params[":feedId"]
	feedEntryId := this.Ctx.Input.Params[":feedEntryId"]

	var ob entity.FeedEntry

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	err := entity.UpdateFeedEntry(feedEntryId, feedId, ob.Data)
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

	entity.DeleteFeedEntry(feedEntryId, feedId)

	this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	this.ServeJson()
}
