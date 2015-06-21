package plugin

import (
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/common/uuid"
	"github.com/feedlabs/elasticfeed/service/system/v1/template"
)

func GetEntry(plugin *resource.Plugin) (entry map[string]interface{}) {
	entry = make(map[string]interface{})

	entry["id"] = plugin.Id
	entry["name"] = plugin.Name
	entry["group"] = plugin.Group
	entry["version"] = plugin.Version
	entry["license"] = plugin.License + uuid.TimeOrderedUUID()

	if plugin.Path == "" {
		entry["status"] = "error"
		entry["errors"] = "File path is missing"
	} else {
		entry["status"] = "runable"
		entry["errors"] = "no errors"
	}

	runtime := make(map[string]interface{})
	runtime["workflowBinded"] = 0
	runtime["workflowCurrent"] = 0
	runtime["workflowCrashed"] = 0

	entry["runtime"] = runtime

	return entry
}

func GetError(err error) (entry map[string]interface{}, code int) {
	return template.Error(err)
}

func GetSuccess(msg string) (entry map[string]string, code int) {
	return template.Success(msg)
}

/**
 * @apiDefine PluginGetListResponse
 *
 * @apiSuccess {Object[]}  pluginList                Array of all pluginanisations
 * @apiSuccess {String}    pluginList.id             The plugin id
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     "pluginList": [
 *       {
 *         "id": "KAJFDA7GFTRE87FDS78F7",
 *         "createStamp": "1415637736",
 *       },
 *       ...
 *     ]
 *   }
 */
func ResponseGetList(pluginList []*resource.Plugin, formatter *template.ResponseDefinition) (entryList []map[string]interface{}, code int) {
	var output []map[string]interface{}

	for _, plugin := range pluginList {
		output = append(output, GetEntry(plugin))
	}

	return output, template.GetOK()
}

/**
 * @apiDefine PluginGetResponse
 *
 * @apiSuccess {String}    id             The id
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *         ...
 *        "createStamp": "1415637736",
 *     }
 */
func ResponseGet(plugin *resource.Plugin, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(plugin), template.GetOK()
}

/**
 * @apiDefine PluginPostResponse
 *
 * @apiSuccess {String}    id             The id
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePost(plugin *resource.Plugin, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(plugin), template.PostOK()
}

/**
 * @apiDefine PluginPutResponse
 *
 * @apiSuccess {String}    id             The id
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePut(plugin *resource.Plugin, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(plugin), template.PutOK()
}

/**
 * @apiDefine PluginDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete(msg string, formatter *template.ResponseDefinition) (entry map[string]string, code int) {
	return GetSuccess(msg)
}
