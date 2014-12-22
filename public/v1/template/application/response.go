package application


/**
 * @apiDefine ApplicationGetListResponse
 *
 * @apiSuccess {Object[]}  applicationList              Array of all applications
 * @apiSuccess {String}    applicationList.id           The application Id
 * @apiSuccess {String}    applicationList.name         The name of the application
 * @apiSuccess {Int}       applicationList.createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "applicationList": [
 *         {
 *           "id": "systemId",
 *           "name": "YourName",
 *           "createStamp": "1234567890",
 *         },
 *         ...
 *       ]
 *     }
 */
func ResponseGetList() {

}

/**
 * @apiDefine ApplicationGetResponse
 *
 * @apiSuccess {String}  id           The application Id
 * @apiSuccess {String}  name         The name of the application
 * @apiSuccess {String}  description  The description of the application
 * @apiSuccess {Int}     createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA786FDS87FDS78F6",
 *       "name": "DragonBall",
 *       "description": "The DragonBall application. This will contain feeds for all DragonBall characters.",
 *       "createStamp": "1415637736",
 *     }
 */
func ResponseGet() {

}

/**
 * @apiDefine ApplicationPostResponse
 *
 * @apiSuccess {String}  id           The application Id
 * @apiSuccess {String}  name         The name of the application
 * @apiSuccess {String}  description  The description of the application
 * @apiSuccess {Int}     createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA786FDS87FDS78F6",
 *       "name": "DragonBall",
 *       "description": "The DragonBall application. This will contain feeds for all DragonBall characters.",
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePost() {

}

/**
 * @apiDefine ApplicationPutResponse
 *
 * @apiSuccess {String}  id           The application Id
 * @apiSuccess {String}  name         The name of the application
 * @apiSuccess {String}  description  The description of the application
 * @apiSuccess {Int}     createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA786FDS87FDS78F6",
 *       "name": "DragonBall",
 *       "description": "The DragonBall application. This will contain feeds for all DragonBall characters.",
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePut() {

}

/**
 * @apiDefine ApplicationDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete() {

}
