package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"cfp/src/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	var ob models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	userid := models.AddUser(ob)
	this.Data["json"] = map[string]string{"UserId": userid}
	this.ServeJson()
}

func (this *UserController) Get() {
	userId := this.Ctx.Input.Params[":objectId"]
	if userId != "" {
		ob, err := models.GetUser(userId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := models.GetUserList()
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *UserController) Put() {
	userId := this.Ctx.Input.Params[":objectId"]
	var ob models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := models.UpdateUser(userId, ob.Name)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJson()
}

func (this *UserController) Delete() {
	userId := this.Ctx.Input.Params[":objectId"]
	models.DeleteUser(userId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}
