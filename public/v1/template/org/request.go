package org

import (
	"errors"
	"github.com/feedlabs/feedify/context"
	"github.com/feedlabs/elasticfeed/public/v1/template"
)

func CheckRequiredParams() {
	// orgId
}

/**
 * @apiDefine OrgGetListRequest
 *
 */
func RequestGetList(input *context.Input) (err error) {
	if template.QueryParamsCount(input.Request.URL) != 0 {
		return errors.New("Too many params in URI query")
	}
	return nil
}

/**
 * @apiDefine OrgGetRequest
 *
 * @apiParam {String} orgId  The org id
 */
func RequestGet(input *context.Input) (err error) {
	if template.QueryParamsCount(input.Request.URL) != 1 {
		return errors.New("Too many params in URI query")
	}
	return nil
}

/**
 * @apiDefine OrgPostRequest
 */
func RequestPost(input *context.Input) (err error) {
	if template.QueryParamsCount(input.Request.URL) != 0 {
		return errors.New("Too many params in URI query")
	}
	return nil
}

/**
 * @apiDefine OrgPutRequest
 *
 * @apiParam {String}    orgId        The org id
 */
func RequestPut(input *context.Input) (err error) {
	if template.QueryParamsCount(input.Request.URL) != 1 {
		return errors.New("Too many params in URI query")
	}
	return nil
}

/**
 * @apiDefine OrgDeleteRequest
 *
 * @apiParam {String}  orgId  The org id
 */
func RequestDelete(input *context.Input) (err error) {
	if template.QueryParamsCount(input.Request.URL) != 1 {
		return errors.New("Too many params in URI query")
	}
	return nil
}
