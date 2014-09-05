package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/feedlabs/api/public/v1/resources"
)

type FeedController struct {
	beego.Controller
}

func (this *FeedController) Post() {
	var ob resources.Feed

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	feedid := resources.AddFeed(ob)
	this.Data["json"] = map[string]string{"id": feedid}
	this.ServeJson()
}

func (this *FeedController) Get() {
	feedId := this.Ctx.Input.Params[":feedId"]
	if feedId != "" {
		ob, err := resources.GetFeed(feedId)
		if err != nil {
			this.Data["json"] = err
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
	resources.DeleteFeed(feedId)

	this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	this.ServeJson()
}
