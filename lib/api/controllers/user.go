package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/feedlabs/feedify/lib/feedify/entity"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	var ob entity.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	userid := entity.AddUser(ob)
	this.Data["json"] = map[string]string{"UserId": userid}
	this.ServeJson()
}

func (this *UserController) Get() {
	userId := this.Ctx.Input.Params[":objectId"]
	if userId != "" {
		ob, err := entity.GetUser(userId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := entity.GetUserList()
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *UserController) Put() {
	userId := this.Ctx.Input.Params[":objectId"]
	var ob entity.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := entity.UpdateUser(userId, ob.Name)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJson()
}

func (this *UserController) Delete() {
	userId := this.Ctx.Input.Params[":objectId"]
	entity.DeleteUser(userId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}
