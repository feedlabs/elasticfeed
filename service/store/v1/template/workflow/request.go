package workflow

import (
	"github.com/feedlabs/feedify/context"
)


/**
 * @apiDefine WorkflowGetListByFeedRequest
 *
 * @apiParam {String}  applicationId  The application id
 * @apiParam {String}  feedId         The application id
 */
func RequestGetList(input *context.Input) {

}

/**
 * @apiDefine WorkflowGetRequest
 *
 * @apiParam {String}  applicationId  The application id
 * @apiParam {String}  workflowId     The workflow id
 */
func RequestGet(input *context.Input) {

}

/**
 * @apiDefine WorkflowPostRequest
 *
 * @apiParam {String}    applicationId  The application id
 * @apiParam {String}    data           The data of the workflow
 * @apiParam {String[]}  [tagList]      Tags of the workflow
 */
func RequestPost(input *context.Input) {

}

/**
 * @apiDefine WorkflowPutRequest
 *
 * @apiParam {String}    applicationId  The application id
 * @apiParam {String}    workflowId     The workflow id
 * @apiParam {String}    data           The data of the workflow
 * @apiParam {String[]}  [tagList]      Tags of the workflow
 */
func RequestPut(input *context.Input) {

}

/**
 * @apiDefine WorkflowDeleteRequest
 *
 * @apiParam {String}  applicationId  The application id
 * @apiParam {String}  workflowId     The workflow id
 */
func RequestDelete(input *context.Input) {

}
