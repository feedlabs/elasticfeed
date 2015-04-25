package workflow

import (
	"errors"
	"github.com/feedlabs/feedify/context"
	"github.com/feedlabs/elasticfeed/service/store/v1/template"
)

func CheckRequiredParams() {
	// workflowId
}

func GetResponseDefinition(input *context.Input) (*template.ResponseDefinition) {
	return template.NewResponseDefinition(input)
}

/**
 * @apiDefine WorkflowGetListRequest
 *
 */
func RequestGetList(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) > 4 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}

/**
 * @apiDefine WorkflowGetRequest
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
 * @apiDefine WorkflowPostRequest
 */
func RequestPost(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) != 0 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}

/**
 * @apiDefine WorkflowPutRequest
 *
 * @apiParam {String}    pluginId        The plugin id
 */
func RequestPut(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) > 4 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}

/**
 * @apiDefine WorkflowDeleteRequest
 *
 * @apiParam {String}  pluginId  The plugin id
 */
func RequestDelete(input *context.Input) (formatter *template.ResponseDefinition, err error) {
	if template.QueryParamsCount(input.Request.URL) != 1 {
		return nil, errors.New("Too many params in URI query")
	}
	return GetResponseDefinition(input), nil
}
