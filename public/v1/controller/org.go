package controller

import (
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
	"github.com/feedlabs/api/public/v1/template/org"
)

type OrgController struct {
	feedify.Controller
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
	org.RequestGetList(this.GetInput())

	obs := resources.GetOrgList()
	this.Data["json"] = obs

	org.ResponseGetList()
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
	org.RequestGet(this.GetInput())
	org.ResponseGet()
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
	org.RequestPost(this.GetInput())

	var ob resources.Org

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	orgid, err := resources.AddOrg(ob)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": orgid}
	}

	org.ResponsePost()
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
	org.RequestPut(this.GetInput())
	org.ResponsePut()
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
	org.RequestDelete(this.GetInput())
	org.ResponseDelete()
	this.ServeJson()
}
