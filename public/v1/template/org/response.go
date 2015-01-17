package org

import (
	"strconv"
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/public/v1/template"
	"errors"
)

func GetEntry(org *resource.Org) (entry map[string]interface{}) {
	entry = make(map[string]interface{})

	stats := make(map[string]string)
	stats["tokens"] = strconv.Itoa(org.Tokens)
	stats["admins"] = strconv.Itoa(org.Admins)
	stats["apps"] = strconv.Itoa(org.Applications)

	entry["id"] = org.Id
	entry["name"] = org.Name
	entry["stats"] = stats

	return entry
}

func GetError(err error) (entry map[string]interface {}, code int) {
	return template.Error(err)
}

func GetSuccess(msg string) (entry map[string]string, code int) {
	return template.Success(msg)
}

/**
 * @apiDefine OrgGetListResponse
 *
 * @apiSuccess {Object[]}  orgList                Array of all organisations
 * @apiSuccess {String}    orgList.id             The org id
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     "orgList": [
 *       {
 *         "id": "KAJFDA7GFTRE87FDS78F7",
 *         "createStamp": "1415637736",
 *       },
 *       ...
 *     ]
 *   }
 */
func ResponseGetList(orgList []*resource.Org, formatter *template.ResponseDefinition) (entryList []map[string]interface{}, code int) {
	var output []map[string]interface{}

	start := formatter.GetPage() * formatter.GetLimit()
	end := start + formatter.GetLimit()
	if len(orgList) < end {
		end = len(orgList)
	}

	if start > end {
		errMsg, _ := GetError(errors.New("Paging is out of range"))
		output = append(output, errMsg)
		return output, template.HTTP_CODE_ACCESS_FORBIDDEN
	}

	orgListPaging := orgList[start:end]

	for _, org := range orgListPaging {
		output = append(output, GetEntry(org))
	}

	return output, template.GetOK()
}

/**
 * @apiDefine OrgGetResponse
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
func ResponseGet(org *resource.Org, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(org), template.GetOK()
}

/**
 * @apiDefine OrgPostResponse
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
func ResponsePost(org *resource.Org, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(org), template.PostOK()
}

/**
 * @apiDefine OrgPutResponse
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
func ResponsePut(org *resource.Org, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(org), template.PutOK()
}

/**
 * @apiDefine OrgDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete(msg string, formatter *template.ResponseDefinition) (entry map[string]string, code int) {
	return GetSuccess(msg)
}
