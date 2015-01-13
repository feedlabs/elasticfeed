package controller

import (
	"encoding/json"

	"github.com/feedlabs/elasticfeed/resource"
	template "github.com/feedlabs/elasticfeed/public/v1/template/org"
)

type OrgController struct {
	DefaultController
}

/**
 * @api {get} org Get List
 * @apiVersion 1.0.0
 * @apiName GetOrgList
 * @apiGroup Organisation
 * @apiDescription This will return a list of all orgs.
 *
 * @apiUse OrgGetListRequest
 * @apiUse OrgGetListResponse
 */
func (this *OrgController) GetList() {
	template.RequestGetList(this.GetInput())

	obs, err := resource.GetOrgList()

	status := 0
	if err != nil {
		this.Data["json"], status = template.GetError(err)
	} else {
		this.Data["json"], status = template.ResponseGetList(obs)
	}

	this.DefaultController.Controller.Ctx.Output.SetStatus(status)

	this.ServeJson()
}

/**
 * @api {get} org/:orgId Get
 * @apiVersion 1.0.0
 * @apiName GetOrg
 * @apiGroup Organisation
 * @apiDescription This will return a specific org.
 *
 * @apiUse OrgGetRequest
 * @apiUse OrgGetResponse
 */
func (this *OrgController) Get() {
	template.RequestGet(this.GetInput())

	orgId := this.Ctx.Input.Params[":orgId"]
	ob, err := resource.GetOrg(orgId)

	status := 0
	if err != nil {
		this.Data["json"], status = template.GetError(err)
	} else {
		this.Data["json"], status = template.ResponseGet(ob)
	}

	this.DefaultController.Controller.Ctx.Output.SetStatus(status)

	this.ServeJson()
}

/**
 * @api {post} org Create
 * @apiVersion 1.0.0
 * @apiName PostOrg
 * @apiGroup Organisation
 * @apiDescription Create a org.
 *
 * @apiUse OrgPostRequest
 * @apiUse OrgPostResponse
 */
func (this *OrgController) Post() {
	template.RequestPost(this.GetInput())

	var org resource.Org

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &org)

	err := resource.AddOrg(&org)

	status := 0
	if err != nil {
		this.Data["json"], status = template.GetError(err)
	} else {
		this.Data["json"], status = template.ResponsePost(&org)
	}

	this.DefaultController.Controller.Ctx.Output.SetStatus(status)

	this.ServeJson()
}

/**
 * @api {put} org/:orgId Update
 * @apiVersion 1.0.0
 * @apiName PutOrg
 * @apiGroup Organisation
 * @apiDescription Update a specific organisation.
 *
 * @apiUse OrgPutRequest
 * @apiUse OrgPutResponse
 */

func (this *OrgController) Put() {
	template.RequestPut(this.GetInput())

	var org resource.Org
	org.Id = this.Ctx.Input.Params[":orgId"]

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &org)

	err := resource.UpdateOrg(&org)

	status := 0
	if err != nil {
		this.Data["json"], status = template.GetError(err)
	} else {
		this.Data["json"], status = template.ResponsePut(&org)
	}

	this.DefaultController.Controller.Ctx.Output.SetStatus(status)

	this.ServeJson()
}

/**
 * @api {delete} org/:orgId Delete
 * @apiVersion 1.0.0
 * @apiName DeleteOrg
 * @apiGroup Organisation
 * @apiDescription Delete a specific org.

 * @apiUse OrgDeleteRequest
 * @apiUse OrgDeleteResponse
 */
func (this *OrgController) Delete() {
	template.RequestDelete(this.GetInput())
	this.DefaultController.Controller.Ctx.Output.SetStatus(template.ResponseDelete())
	this.ServeJson()
}
