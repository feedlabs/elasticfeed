package org

import (
	"strconv"
	"github.com/feedlabs/api/resource"
)

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
func ResponseGetList(orgList []*resource.Org) []map[string]interface {} {

	var output []map[string]interface {}

	for _, value := range orgList {
		m := make(map[string]interface {})

		stats := make(map[string]string)
		stats["tokens"] = strconv.Itoa(value.Tokens)
		stats["admins"] = strconv.Itoa(value.Admins)
		stats["apps"] = strconv.Itoa(value.Applications)

		m["id"] = value.Id
		m["key"] = value.ApiKey
		m["rel"] = stats

		output = append(output, m)
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
func ResponseGet() {

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
