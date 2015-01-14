package org

import (
	"errors"
	"github.com/feedlabs/feedify/context"
	"github.com/feedlabs/elasticfeed/public/v1/template"
)

/**
 * @apiDefine OrgGetListRequest
 *
 */
func RequestGetList(input *context.Input) (err error) {
	if template.HasQueryParams(input.Request.URL) {
		return errors.New("Too many params in URI query")
	}
	return nil
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
