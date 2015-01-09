package token

/**
 * @apiDefine OrgTokenGetListResponse
 *
 * @apiSuccess {Object[]}  orgTokenList                Array of all organisation tokens
 * @apiSuccess {String}    orgTokenList.id             The orgToken id
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     "orgTokenList": [
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
 * @apiDefine OrgTokenGetResponse
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
 * @apiDefine OrgTokenPostResponse
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
 * @apiDefine OrgTokenPutResponse
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
 * @apiDefine OrgTokenDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete() {

}
