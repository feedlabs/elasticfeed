package plugin

import (
	"errors"
	"github.com/feedlabs/feedify/context"
	"github.com/feedlabs/elasticfeed/service/system/template"
)

func CheckRequiredParams() {
	// pluginId
}

func GetResponseDefinition(input *context.Input) (*template.ResponseDefinition) {
	return template.NewResponseDefinition(input)
}

/**
 * @apiDefine OrgGetListRequest
 *
 */
func RequestGetList(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) > 4 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}

/**
 * @apiDefine OrgGetRequest
 *
 * @apiParam {String} pluginId  The plugin id
 */
func RequestGet(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) != 1 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}

/**
 * @apiDefine OrgPostRequest
 */
func RequestPost(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) != 0 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}

/**
 * @apiDefine OrgPutRequest
 *
 * @apiParam {String}    pluginId        The plugin id
 */
func RequestPut(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) != 1 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}

/**
 * @apiDefine OrgDeleteRequest
 *
 * @apiParam {String}  pluginId  The plugin id
 */
func RequestDelete(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) != 1 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}
