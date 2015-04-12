package controller

import (
	"encoding/json"

	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/service/store/v1/template/application"
)

type ApplicationController struct {
	DefaultController
}

/**
 * @api {get} application Get List
 * @apiVersion 1.0.0
 * @apiName GetApplicationList
 * @apiGroup Application
 * @apiDescription This will return a list of all applications you have created.
 *
 * @apiUse ApplicationGetListRequest
 * @apiUse ApplicationGetListResponse
 */
func (this *ApplicationController) GetList() {
	application.RequestGetList(this.GetInput())

	obs, err := resource.GetApplicationList(this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	application.ResponseGetList()
	this.Controller.ServeJson()
}

/**
 * @api {get} application/:applicationId Get
 * @apiVersion 1.0.0
 * @apiName GetApplication
 * @apiGroup Application
 * @apiDescription This will return a specific application.
 *
 * @apiUse ApplicationGetRequest
 * @apiUse ApplicationGetResponse
 */
func (this *ApplicationController) Get() {
	application.RequestGet(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	ob, err := resource.GetApplication(appId, this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = ob
	}

	application.ResponseGet()
	this.Controller.ServeJson()
}

/**
 * @api {post} application Create
 * @apiVersion 1.0.0
 * @apiName PostApplication
 * @apiGroup Application
 * @apiDescription Create a application.
 *
 * @apiUse ApplicationPostRequest
 * @apiUse ApplicationPostResponse
 */
func (this *ApplicationController) Post() {
	application.RequestPost(this.GetInput())

	var ob resource.Application

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	appid, err := resource.AddApplication(ob, this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": appid}
	}

	application.ResponsePost()
	this.Controller.ServeJson()
}

/**
 * @api {put} application/:applicationId Update
 * @apiVersion 1.0.0
 * @apiName PutApplication
 * @apiGroup Application
 * @apiDescription Update a specific application.
 *
 * @apiUse ApplicationPutRequest
 * @apiUse ApplicationPutResponse
 */

func (this *ApplicationController) Put() {
	application.RequestPut(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	var ob resource.Application

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	err := resource.UpdateApplication(appId, ob.Data)
	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "update success", "status": "ok"}
	}

	application.ResponsePut()
	this.Controller.ServeJson()
}

/**
 * @api {delete} application/:applicationId Delete
 * @apiVersion 1.0.0
 * @apiName DeleteApplication
 * @apiGroup Application
 * @apiDescription Delete a specific application.
 *
 * @apiUse ApplicationDeleteRequest
 * @apiUse ApplicationDeleteResponse
 */
func (this *ApplicationController) Delete() {
	application.RequestDelete(this.GetInput())

	appId := this.Ctx.Input.Params[":applicationId"]
	err := resource.DeleteApplication(appId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"result": "delete success", "status": "ok"}
	}

	application.ResponseDelete()
	this.Controller.ServeJson()
}
