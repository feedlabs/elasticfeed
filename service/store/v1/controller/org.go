package controller

import (
	"encoding/json"

	"github.com/feedlabs/elasticfeed/resource"
	template "github.com/feedlabs/elasticfeed/service/store/v1/template/org"
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
	formatter, err := template.RequestGetList(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	obs, err := resource.GetOrgList()
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponseGetList(obs, formatter))
	}
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
	formatter, err := template.RequestGet(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	orgId := this.Ctx.Input.Params[":orgId"]

	ob, err := resource.GetOrg(orgId)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponseGet(ob, formatter))
	}
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
	formatter, err := template.RequestPost(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	var org resource.Org

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &org)

	err = resource.AddOrg(&org)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponsePost(&org, formatter))
	}
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
	formatter, err := template.RequestPut(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	var org resource.Org
	org.Id = this.Ctx.Input.Params[":orgId"]

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &org)

	err = resource.UpdateOrg(&org)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponsePut(&org, formatter))
	}
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
	formatter, err := template.RequestDelete(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	} else {
		this.ServeJson(template.ResponseDelete("Org has been deleted", formatter))
	}
}
