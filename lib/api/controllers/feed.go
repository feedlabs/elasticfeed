package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/feedlabs/feedify/lib/feedify/entity"
)

type FeedController struct {
	beego.Controller
}

func (this *FeedController) Post() {
	var ob entity.Feed

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	feedid := entity.AddFeed(ob)
	this.Data["json"] = map[string]string{"id": feedid}
	this.ServeJson()
}

func (this *FeedController) Get() {
	feedId := this.Ctx.Input.Params[":feedId"]
	if feedId != "" {
		ob, err := entity.GetFeed(feedId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := entity.GetFeedList()
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *FeedController) Put() {
	feedId := this.Ctx.Input.Params[":feedId"]
	var ob entity.Feed

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	err := entity.UpdateFeed(feedId, ob.Data)
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "update success", "status": "ok"}
	}
	this.ServeJson()
}

func (this *FeedController) Delete() {
	feedId := this.Ctx.Input.Params[":feedId"]
	entity.DeleteFeed(feedId)

	this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	this.ServeJson()
}
