package org

import (
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
	"github.com/feedlabs/api/public/v1/template/org/token"
)

type TokenController struct {
	feedify.Controller
}

/**
 * @api {get} org/:orgId/token Get List (Org)
 * @apiVersion 1.0.0
 * @apiName GetOrgTokenList
 * @apiGroup Token
 * @apiDescription This will return a list of all organisation tokens
 *
 * @apiUse OrgTokenGetListRequest
 * @apiUse OrgTokenGetListResponse
 */
func (this *TokenController) GetList() {
	token.RequestGetList(this.GetInput())

	orgId := this.Ctx.Input.Params[":orgId"]
	obs, err := resources.GetOrgTokenList(orgId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	token.ResponseGetList()
	this.ServeJson()
}

/**
 * @api {get} org/:orgId/token/:tokenId Get (Org)
 * @apiVersion 1.0.0
 * @apiName GetOrgToken
 * @apiGroup Token
 * @apiDescription This will return a specific organisation token
 *
 * @apiUse OrgTokenGetRequest
 * @apiUse OrgTokenGetResponse
 */
func (this *TokenController) Get() {
	token.RequestGet(this.GetInput())

	orgId := this.Ctx.Input.Params[":orgId"]
	tokenId := this.Ctx.Input.Params[":tokenId"]
	ob, err := resources.GetOrgToken(tokenId, orgId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = ob
	}

	token.ResponseGet()
	this.ServeJson()
}

/**
 * @api {post} org/:orgId/token Create (Org)
 * @apiVersion 1.0.0
 * @apiName PostOrgToken
 * @apiGroup Token
 * @apiDescription Create a organisation token
 *
 * @apiUse OrgTokenPostRequest
 * @apiUse OrgTokenPostResponse
 */
func (this *TokenController) Post() {
	token.RequestPost(this.GetInput())

	var ob resources.Token

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	orgId := this.Ctx.Input.Params[":orgId"]
	tokenid, err := resources.AddOrgToken(ob, orgId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": tokenid}
	}

	token.ResponsePost()
	this.ServeJson()
}

/**
 * @api {delete} org/:orgId/token/:tokenId Delete (Org)
 * @apiVersion 1.0.0
 * @apiName DeleteOrgToken
 * @apiGroup Token
 * @apiDescription Delete a specific organisation token

 * @apiUse OrgTokenDeleteRequest
 * @apiUse OrgTokenDeleteResponse
 */
func (this *TokenController) Delete() {
	token.RequestDelete(this.GetInput())

//	orgId := this.Ctx.Input.Params[":orgId"]
//	tokenId := this.Ctx.Input.Params[":tokenId"]

	token.ResponseDelete()
	this.ServeJson()
}
