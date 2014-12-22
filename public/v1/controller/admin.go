package controller

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/public/v1/template/admin"
)

type AdminController struct {
	feedify.Controller
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
	admin.ResponseGetList()
	this.ServeJson()
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
	admin.ResponseGet()
	this.ServeJson()
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
	admin.ResponsePost()
	this.ServeJson()
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
	this.ServeJson()
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
	this.ServeJson()
}
