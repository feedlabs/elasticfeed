package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/feedlabs/cfp/models"
)

type StorageController struct {
	beego.Controller
}

func (this *StorageController) Post() {
	var ob models.Storage
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	storageid := models.AddStorage(ob)
	this.Data["json"] = map[string]string{"StorageId": storageid}
	this.ServeJson()
}

func (this *StorageController) Get() {
	storageId := this.Ctx.Input.Param(":id")
	if storageId != "" {
		ob, err := models.GetStorage(storageId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := models.GetStorageList()
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *StorageController) Put() {
	storageId := this.Ctx.Input.Param(":id")
	var ob models.Storage
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := models.UpdateStorage(storageId, ob.Name)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJson()
}

func (this *StorageController) Delete() {
	storageId := this.Ctx.Input.Param(":id")
	models.DeleteStorage(storageId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}

func (this *StorageController) Options() {
	this.Data["json"] = "GET"
	this.ServeJson()
}
