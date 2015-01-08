package controller

import (
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
	"github.com/feedlabs/api/public/v1/template/token"
)

type TokenController struct {
	feedify.Controller
}

/**
 * @api {get} token Get List
 * @apiVersion 1.0.0
 * @apiName GetTokenList
 * @apiGroup Token
 * @apiDescription This will return a list of all tokens.
 *
 * @apiUse TokenGetListRequest
 * @apiUse TokenGetListResponse
 */
func (this *TokenController) GetList() {
	token.RequestGetList(this.GetInput())

	adminId := this.Ctx.Input.Params[":adminId"]
	obs, err := resources.GetTokenList(adminId, GetMyOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	token.ResponseGetList()
	this.ServeJson()
}

/**
 * @api {get} token/:token Get
 * @apiVersion 1.0.0
 * @apiName GetToken
 * @apiGroup Token
 * @apiDescription This will return a specific token.
 *
 * @apiUse TokenGetRequest
 * @apiUse TokenGetResponse
 */
func (this *TokenController) Get() {
	token.RequestGet(this.GetInput())
	token.ResponseGet()
	this.ServeJson()
}

/**
 * @api {post} token Create
 * @apiVersion 1.0.0
 * @apiName PostToken
 * @apiGroup Token
 * @apiDescription Create a token.
 *
 * @apiUse TokenPostRequest
 * @apiUse TokenPostResponse
 */
func (this *TokenController) Post() {
	token.RequestPost(this.GetInput())

	var ob resources.Token

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	adminId := this.Ctx.Input.Params[":adminId"]
	appid, err := resources.AddToken(ob, adminId, GetMyOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": appid}
	}

	token.ResponsePost()
	this.ServeJson()
}

/**
 * @api {delete} token/:token Delete
 * @apiVersion 1.0.0
 * @apiName DeleteToken
 * @apiGroup Token
 * @apiDescription Delete a specific token.

 * @apiUse TokenDeleteRequest
 * @apiUse TokenDeleteResponse
 */
func (this *TokenController) Delete() {
	token.RequestDelete(this.GetInput())
	token.ResponseDelete()
	this.ServeJson()
}
