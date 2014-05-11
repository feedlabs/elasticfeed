package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/feedlabs/feedify/models"
)

type FeedController struct {
	beego.Controller
}

func (this *FeedController) Post() {
	var ob models.Feed
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	feedid := models.AddFeed(ob)
	this.Data["json"] = map[string]string{"FeedId": feedid}
	this.ServeJson()
}

func (this *FeedController) Get() {
	feedId := this.Ctx.Input.Params[":objectId"]
	if feedId != "" {
		ob, err := models.GetFeed(feedId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := models.GetFeedList()
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *FeedController) Put() {
	feedId := this.Ctx.Input.Params[":objectId"]
	var ob models.Feed
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := models.UpdateFeed(feedId, ob.Name)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJson()
}

func (this *FeedController) Delete() {
	feedId := this.Ctx.Input.Params[":objectId"]
	models.DeleteFeed(feedId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}
