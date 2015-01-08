package admin

import (
	"encoding/json"

	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/resources"
	"github.com/feedlabs/api/public/v1/controller"
	"github.com/feedlabs/api/public/v1/template/admin/token"
)

type TokenController struct {
	feedify.Controller
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
func (this *TokenController) GetList() {
	token.RequestGetList(this.GetInput())

	adminId := this.Ctx.Input.Params[":adminId"]
	obs, err := resources.GetTokenList(adminId, controller.GetMyOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = obs
	}

	token.ResponseGetList()
	this.ServeJson()
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
func (this *TokenController) Get() {
	token.RequestGet(this.GetInput())
	token.ResponseGet()
	this.ServeJson()
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
func (this *TokenController) Post() {
	token.RequestPost(this.GetInput())

	var ob resources.Token

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &ob)

	adminId := this.Ctx.Input.Params[":adminId"]
	appid, err := resources.AddToken(ob, adminId, controller.GetMyOrgId())

	if err != nil {
		this.Data["json"] = map[string]string{"result": err.Error(), "status": "error"}
	} else {
		this.Data["json"] = map[string]string{"id": appid}
	}

	token.ResponsePost()
	this.ServeJson()
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
func (this *TokenController) Delete() {
	token.RequestDelete(this.GetInput())
	token.ResponseDelete()
	this.ServeJson()
}
