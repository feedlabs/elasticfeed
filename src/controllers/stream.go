package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/feedlabs/cfp/models"
)

type StreamController struct {
	beego.Controller
}

func (this *StreamController) Post() {
	var ob models.Stream
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	streamid := models.AddStream(ob)
	this.Data["json"] = map[string]string{"StreamId": streamid}
	this.ServeJson()
}

func (this *StreamController) Get() {
	streamId := this.Ctx.Input.Params[":objectId"]
	if streamId != "" {
		ob, err := models.GetStream(streamId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := models.GetStreamList()
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *StreamController) Put() {
	streamId := this.Ctx.Input.Params[":objectId"]
	var ob models.Stream
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := models.UpdateStream(streamId, ob.Name)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJson()
}

func (this *StreamController) Delete() {
	streamId := this.Ctx.Input.Params[":objectId"]
	models.DeleteStream(streamId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}
