package org

import (
	"strconv"
	"github.com/feedlabs/elasticfeed/resource"
)

func GetEntry(org *resource.Org) (entry map[string]interface{}) {
	entry = make(map[string]interface{})

	stats := make(map[string]string)
	stats["tokens"] = strconv.Itoa(org.Tokens)
	stats["admins"] = strconv.Itoa(org.Admins)
	stats["apps"] = strconv.Itoa(org.Applications)

	entry["id"] = org.Id
	entry["key"] = org.ApiKey
	entry["stats"] = stats

	return entry
}

func GetError(err error) map[string]string {
	return map[string]string{"result": err.Error(), "status": "error"}
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
func ResponseGetList(orgList []*resource.Org) []map[string]interface{} {
	var output []map[string]interface{}

	for _, org := range orgList {
		output = append(output, GetEntry(org))
	}

	return output
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
func ResponseGet(org *resource.Org) map[string]interface{} {
	return GetEntry(org)
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
func ResponsePost() {

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
func ResponsePut() {

}

/**
 * @apiDefine OrgDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete() {

}
