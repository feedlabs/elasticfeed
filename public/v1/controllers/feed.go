package controllers

import (
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
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

func (this *FeedController) Get() {
	feedId := this.Ctx.Input.Params[":feedId"]
	if feedId != "" {
		ob, err := resources.GetFeed(feedId)

		if err != nil {
			this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := resources.GetFeedList()
		this.Data["json"] = obs
	}
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

	result := "delete success"
	if err != nil {
		result = err.Error()
	}

	this.Data["json"] = map[string]string{"result": result, "status": "ok"}
	this.ServeJson()
}
