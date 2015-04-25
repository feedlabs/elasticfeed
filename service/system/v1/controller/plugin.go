package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/feedlabs/elasticfeed/common/config"
	"github.com/feedlabs/elasticfeed/common/uuid"
	"github.com/feedlabs/elasticfeed/resource"
	template "github.com/feedlabs/elasticfeed/service/system/v1/template/plugin"
)

type PluginController struct {
	DefaultController
}

/**
 * @api {get} plugin Get List
 * @apiVersion 1.0.0
 * @apiName GetPluginList
 * @apiGroup Plugin
 * @apiDescription This will return a list of all plugins.
 *
 * @apiUse PluginGetListRequest
 * @apiUse PluginGetListResponse
 */
func (this *PluginController) GetList() {
	formatter, err := template.RequestGetList(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	obs, err := resource.GetPluginList()
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponseGetList(obs, formatter))
	}
}

/**
 * @api {get} plugin/:pluginId Get
 * @apiVersion 1.0.0
 * @apiName GetPlugin
 * @apiGroup Plugin
 * @apiDescription This will return a specific plugin.
 *
 * @apiUse PluginGetRequest
 * @apiUse PluginGetResponse
 */
func (this *PluginController) Get() {
	formatter, err := template.RequestGet(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	pluginId := this.Ctx.Input.Params[":pluginId"]

	ob, err := resource.GetPlugin(pluginId)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponseGet(ob, formatter))
	}
}

/**
 * @api {post} plugin Create
 * @apiVersion 1.0.0
 * @apiName PostPlugin
 * @apiGroup Plugin
 * @apiDescription Create a plugin.
 *
 * @apiUse PluginPostRequest
 * @apiUse PluginPostResponse
 */
func (this *PluginController) Post() {
	formatter, err := template.RequestPost(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	var plugin resource.Plugin

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &plugin)

	err = resource.AddPlugin(&plugin)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponsePost(&plugin, formatter))
	}
}

/**
 * @api {put} plugin/:pluginId Update
 * @apiVersion 1.0.0
 * @apiName PutPlugin
 * @apiGroup Plugin
 * @apiDescription Update a specific plugin.
 *
 * @apiUse PluginPutRequest
 * @apiUse PluginPutResponse
 */

func (this *PluginController) Put() {
	formatter, err := template.RequestPut(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	var plugin resource.Plugin
	plugin.Id = this.Ctx.Input.Params[":pluginId"]

	data := this.Ctx.Input.CopyBody()
	json.Unmarshal(data, &plugin)

	err = resource.UpdatePlugin(&plugin)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponsePut(&plugin, formatter))
	}
}

/**
 * @api {put} plugin/:pluginId/upload Upload
 * @apiVersion 1.0.0
 * @apiName PutPluginFile
 * @apiGroup Plugin
 * @apiDescription Update a specific plugin.
 *
 * @apiUse PluginPutFileRequest
 * @apiUse PluginPutFileResponse
 */
func (this *PluginController) PutFile() {

	formatter, err := template.RequestPut(this.GetInput())
	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	}

	pluginId := this.Ctx.Input.Params[":pluginId"]
	plugin, err := resource.GetPlugin(pluginId)

	rel_path := config.GetPluginStoragePath() + "/" + plugin.Group + "-" + uuid.TimeOrderedUUID() + "-" + pluginId
	abs_path := config.GetHomeAbsolutePath() + "/" + rel_path

	data := this.Ctx.Input.CopyBody()
	err = ioutil.WriteFile(abs_path, data, 0644)
	if err != nil {
		panic(err)
	}

	plugin.Path = rel_path

	err = resource.UpdatePlugin(plugin)
	if err != nil {
		this.ServeJson(template.GetError(err))
	} else {
		this.ServeJson(template.ResponsePut(plugin, formatter))
	}
}

/**
 * @api {delete} plugin/:pluginId Delete
 * @apiVersion 1.0.0
 * @apiName DeletePlugin
 * @apiGroup Plugin
 * @apiDescription Delete a specific plugin.

 * @apiUse PluginDeleteRequest
 * @apiUse PluginDeleteResponse
 */
func (this *PluginController) Delete() {
	formatter, err := template.RequestDelete(this.GetInput())

	pluginId := this.Ctx.Input.Params[":pluginId"]

	err = resource.DeletePlugin(pluginId)

	if err != nil {
		this.ServeJson(template.GetError(err))
		return
	} else {
		this.ServeJson(template.ResponseDelete("Plugin has been deleted", formatter))
	}
}
