package controller

import (
	"encoding/json"

	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/service/store/v1/template/admin"
)

type AdminController struct {
	DefaultController
}

/**
 * @api {get} admin Get List
 * @apiVersion 1.0.0
 * @apiName GetAdminList
 * @apiGroup Admin
 * @apiDescription This will return a list of all admin users.
 *
 * @apiUse AdminGetListRequest
 * @apiUse AdminGetListResponse
 */
func (this *AdminController) GetList() {
	admin.RequestGetList(this.GetInput())

	obs, err := resource.GetAdminList(this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	admin.ResponseGetList()
	this.Controller.ServeJson()
}

/**
 * @api {get} admin/:adminId Get
 * @apiVersion 1.0.0
 * @apiName GetAdmin
 * @apiGroup Admin
 * @apiDescription This will return a specific token.
 *
 * @apiUse AdminGetRequest
 * @apiUse AdminGetResponse
 */
func (this *AdminController) Get() {
	admin.RequestGet(this.GetInput())

	adminId := this.Ctx.Input.Params[":adminId"]
	ob, err := resource.GetAdmin(adminId, this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = ob
	}

	admin.ResponseGet()
	this.Controller.ServeJson()
}

/**
 * @api {post} admin Create
 * @apiVersion 1.0.0
 * @apiName PostAdmin
 * @apiGroup Admin
 * @apiDescription Create a admin user.
 *
 * @apiUse AdminPostRequest
 * @apiUse AdminPostResponse
 */
func (this *AdminController) Post() {
	admin.RequestPost(this.GetInput())

	var ob resource.Admin

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	adminid, err := resource.AddAdmin(ob, this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": adminid}
	}

	admin.ResponsePost()
	this.Controller.ServeJson()
}

/**
 * @api {put} admin/:adminId Update
 * @apiVersion 1.0.0
 * @apiName PutAdmin
 * @apiGroup Admin
 * @apiDescription Update a specific admin user.
 *
 * @apiUse AdminPutRequest
 * @apiUse AdminPutResponse
 */
func (this *AdminController) Put() {
	admin.RequestPut(this.GetInput())
	admin.ResponsePut()
	this.Controller.ServeJson()
}

/**
 * @api {delete} admin/:adminId Delete
 * @apiVersion 1.0.0
 * @apiName DeleteAdmin
 * @apiGroup Admin
 * @apiDescription Delete a specific admin user.
 *
 * @apiUse AdminDeleteRequest
 * @apiUse AdminDeleteResponse
 */
func (this *AdminController) Delete() {
	admin.RequestDelete(this.GetInput())
	admin.ResponseDelete()
	this.Controller.ServeJson()
}
