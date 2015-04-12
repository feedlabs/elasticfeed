package controller

import (
	"encoding/json"

	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/service/store/v1/template/token"
)

type TokenController struct {
	DefaultController
}

/**
 * @api {get} org/:orgId/token Get List (Org)
 * @apiVersion 1.0.0
 * @apiName GetOrgTokenList
 * @apiGroup Token
 * @apiDescription This will return a list of all organisation tokens
 *
 * @apiUse TokenGetListRequest
 * @apiUse TokenGetListResponse
 */
func (this *TokenController) GetOrgList() {
	token.RequestGetList(this.GetInput())

	orgId := this.Ctx.Input.Params[":orgId"]
	obs, err := resource.GetOrgTokenList(orgId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	token.ResponseGetList()
	this.Controller.ServeJson()
}

/**
 * @api {get} org/:orgId/token/:tokenId Get (Org)
 * @apiVersion 1.0.0
 * @apiName GetOrgToken
 * @apiGroup Token
 * @apiDescription This will return a specific organisation token
 *
 * @apiUse TokenGetRequest
 * @apiUse TokenGetResponse
 */
func (this *TokenController) GetOrg() {
	token.RequestGet(this.GetInput())

	orgId := this.Ctx.Input.Params[":orgId"]
	tokenId := this.Ctx.Input.Params[":tokenId"]
	ob, err := resource.GetOrgToken(tokenId, orgId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = ob
	}

	token.ResponseGet()
	this.Controller.ServeJson()
}

/**
 * @api {post} org/:orgId/token Create (Org)
 * @apiVersion 1.0.0
 * @apiName PostOrgToken
 * @apiGroup Token
 * @apiDescription Create a organisation token
 *
 * @apiUse TokenPostRequest
 * @apiUse TokenPostResponse
 */
func (this *TokenController) PostOrg() {
	token.RequestPost(this.GetInput())

	var ob resource.Token

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	orgId := this.Ctx.Input.Params[":orgId"]
	tokenid, err := resource.AddOrgToken(ob, orgId)

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": tokenid}
	}

	token.ResponsePost()
	this.Controller.ServeJson()
}

/**
 * @api {delete} org/:orgId/token/:tokenId Delete (Org)
 * @apiVersion 1.0.0
 * @apiName DeleteOrgToken
 * @apiGroup Token
 * @apiDescription Delete a specific organisation token

 * @apiUse TokenDeleteRequest
 * @apiUse TokenDeleteResponse
 */
func (this *TokenController) DeleteOrg() {
	token.RequestDelete(this.GetInput())

//	orgId := this.Ctx.Input.Params[":orgId"]
//	tokenId := this.Ctx.Input.Params[":tokenId"]

	token.ResponseDelete()
	this.Controller.ServeJson()
}

/**
 * @api {get} admin/:adminId/token Get List (Admin)
 * @apiVersion 1.0.0
 * @apiName GetTokenList
 * @apiGroup Token
 * @apiDescription This will return a list of all tokens for specific admin.
 *
 * @apiUse TokenGetListRequest
 * @apiUse TokenGetListResponse
 */
func (this *TokenController) GetAdminList() {
	token.RequestGetList(this.GetInput())

	adminId := this.Ctx.Input.Params[":adminId"]
	obs, err := resource.GetTokenList(adminId, this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	token.ResponseGetList()
	this.Controller.ServeJson()
}

/**
 * @api {get} admin/:adminId/token/:token Get (Admin)
 * @apiVersion 1.0.0
 * @apiName GetToken
 * @apiGroup Token
 * @apiDescription This will return a specific token for specific admin.
 *
 * @apiUse TokenGetRequest
 * @apiUse TokenGetResponse
 */
func (this *TokenController) GetAdmin() {
	token.RequestGet(this.GetInput())
	token.ResponseGet()
	this.Controller.ServeJson()
}

/**
 * @api {post} admin/:adminId/token Create (Admin)
 * @apiVersion 1.0.0
 * @apiName PostToken
 * @apiGroup Token
 * @apiDescription Create a token for specific admin.
 *
 * @apiUse TokenPostRequest
 * @apiUse TokenPostResponse
 */
func (this *TokenController) PostAdmin() {
	token.RequestPost(this.GetInput())

	var ob resource.Token

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	adminId := this.Ctx.Input.Params[":adminId"]
	appid, err := resource.AddToken(ob, adminId, this.GetAdminOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": appid}
	}

	token.ResponsePost()
	this.Controller.ServeJson()
}

/**
 * @api {delete} admin/:adminId/token/:token Delete (Admin)
 * @apiVersion 1.0.0
 * @apiName DeleteToken
 * @apiGroup Token
 * @apiDescription Delete a specific token for specific admin.

 * @apiUse TokenDeleteRequest
 * @apiUse TokenDeleteResponse
 */
func (this *TokenController) DeleteAdmin() {
	token.RequestDelete(this.GetInput())
	token.ResponseDelete()
	this.Controller.ServeJson()
}
