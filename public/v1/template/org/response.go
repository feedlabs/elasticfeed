package org

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
func ResponseGetList() {

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
