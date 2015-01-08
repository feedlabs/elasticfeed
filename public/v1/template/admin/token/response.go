package token


/**
 * @apiDefine TokenGetListResponse
 *
 * @apiSuccess {Object[]}  tokenList              Array of all tokens
 * @apiSuccess {String}    tokenList.token        The token
 * @apiSuccess {String}    tokenList.name         The name of the token
 * @apiSuccess {Int}       tokenList.createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     "tokenList": [
 *       {
 *         "token": "KAJFDA7GFTRE87FDS78F7",
 *         "name": "Super Saiyan",
 *         "createStamp": "1415637736",
 *       },
 *       ...
 *     ]
 *   }
 */
func ResponseGetList() {

}

/**
 * @apiDefine TokenGetResponse
 *
 * @apiSuccess {String}    token        The token
 * @apiSuccess {String}    name         The name of the token
 * @apiSuccess {Int}       createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "token": "KAJFDA7GFTRE87FDS78F7",
 *       "name": "Super Saiyan",
 *       "createStamp": "1415637736",
 *     }
 */
func ResponseGet() {

}

/**
 * @apiDefine TokenPostResponse
 *
 * @apiSuccess {String}    token        The token
 * @apiSuccess {String}    name         The name of the token
 * @apiSuccess {Int}       createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "token": "KAJFDA7GFTRE87FDS78F7",
 *       "name": "Super Saiyan",
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePost() {

}

/**
 * @apiDefine TokenDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete() {

}
