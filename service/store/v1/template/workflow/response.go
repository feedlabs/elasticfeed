package workflow

import (
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/service/store/v1/template"
)

func GetEntry(workflow *resource.Workflow) (entry map[string]interface{}) {
	entry = make(map[string]interface{})

	entry["id"] = workflow.Id
	entry["applicationId"] = workflow.Feed.Application.Id
	entry["feedId"] = workflow.Feed.Id
	entry["default"] = workflow.Default
	entry["data"] = workflow.Data
	entry["status"] = "running"
	entry["errors"] = "no errors"

	return entry
}

func GetError(err error) (entry map[string]interface{}, code int) {
	return template.Error(err)
}

func GetSuccess(msg string) (entry map[string]string, code int) {
	return template.Success(msg)
}
/**
 * @apiDefine WorkflowGetListByFeedResponse
 *
 * @apiSuccess {Object}    feed
 * @apiSuccess {String}    feed.id                The application Id
 * @apiSuccess {String}    feed.name              The name of the application
 * @apiSuccess {String}    feed.channelId         The channel Id of the feed
 * @apiSuccess {String}    feed.applicationId     The id of the application the feed belongs to
 * @apiSuccess {Int}       feed.createStamp       Unix time stamp of create time
 * @apiSuccess {Object[]}  workflowList              Array of all workflows
 * @apiSuccess {String}    workflowList.id           The workflow Id
 * @apiSuccess {Int}       workflowList.createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "feed": {
 *           "id": "KAJFDA7GFTRE87FDS78F7",
 *           "name": "Son Goku",
 *           "channelId": "ASJDH86ASD678ASDASD768",
 *           "applicationId": "KAJFDA786FDS87FDS78F6",
 *           "createStamp": "1415637736",
 *       }
 *       "workflowList": [
 *         {
 *           "id": "KAJFDA7GFTRE87FDS78F7",
 *           "createStamp": "1415637736",
 *         },
 *         ...
 *       ]
 *     }
 */
func ResponseGetList(workflowList []*resource.Workflow, formatter *template.ResponseDefinition) (entryList []map[string]interface{}, code int) {
	var output []map[string]interface{}

	for _, plugin := range workflowList {
		output = append(output, GetEntry(plugin))
	}

	return output, template.GetOK()
}

/**
 * @apiDefine WorkflowGetResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    data           The data of the workflow
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponseGet(workflow *resource.Workflow, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(workflow), template.GetOK()
}

/**
 * @apiDefine WorkflowPostResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    data           The data of the workflow
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePost(workflow *resource.Workflow, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(workflow), template.PostOK()
}

/**
 * @apiDefine WorkflowPutResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    [feedId]      The feed id
 * @apiSuccess {String}    data           The data of the workflow
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "feedId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePut(workflow *resource.Workflow, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(workflow), template.PutOK()
}

/**
 * @apiDefine WorkflowDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete(msg string, formatter *template.ResponseDefinition) (entry map[string]string, code int) {
	return GetSuccess(msg)
}

