package controller

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/public/v1/template/application"
)

type ApplicationController struct {
	feedify.Controller
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
	application.ResponseGetList()
	this.ServeJson()
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
	application.ResponseGet()
	this.ServeJson()
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
	application.ResponsePost()
	this.ServeJson()
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
	application.ResponsePut()
	this.ServeJson()
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
	application.ResponseDelete()
	this.ServeJson()
}
