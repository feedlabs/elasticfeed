package org

import (
//	"fmt"
//	"errors"
	"github.com/feedlabs/feedify/context"
)


/**
 * @apiDefine OrgGetListRequest
 *
 */
func RequestGetList(input *context.Input) (err error) {
	return nil //errors.New("worng request")
}

/**
 * @apiDefine OrgGetRequest
 *
 * @apiParam {String} orgId  The org id
 */
func RequestGet(input *context.Input) {

}

/**
 * @apiDefine OrgPostRequest
 */
func RequestPost(input *context.Input) {

}

/**
 * @apiDefine OrgPutRequest
 *
 * @apiParam {String}    orgId        The org id
 */
func RequestPut(input *context.Input) {

}

/**
 * @apiDefine OrgDeleteRequest
 *
 * @apiParam {String}  orgId  The org id
 */
func RequestDelete(input *context.Input) {

}
