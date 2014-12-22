package feed


/**
 * @apiDefine FeedGetListResponse
 *
 * @apiSuccess {Object}    application
 * @apiSuccess {String}    application.id           The application Id
 * @apiSuccess {String}    application.name         The name of the application
 * @apiSuccess {Int}       application.createStamp  Unix time stamp of create time
 * @apiSuccess {Object[]}  feedList                 Array of all feeds
 * @apiSuccess {String}    feedList.id              The feed Id
 * @apiSuccess {String}    feedList.name            The name of the feed
 * @apiSuccess {String}    feedList.channelId       The channel Id of the feed
 * @apiSuccess {Int}       feedList.createStamp     Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "application": {
 *           "id": "KAJFDA786FDS87FDS78F6",
 *           "name": "DragonBall",
 *           "createStamp": "1415637736",
 *       }
 *       "feedList": [
 *         {
 *           "id": "KAJFDA7GFTRE87FDS78F7",
 *           "name": "Son Goku",
 *           "channelId": "ASJDH86ASD678ASDASD768",
 *           "createStamp": "1415637736",
 *         },
 *         ...
 *       ]
 *     }
 */
func ResponseGetList() {

}

/**
 * @apiDefine FeedGetResponse
 *
 * @apiSuccess {String}    id             The feed id
 * @apiSuccess {String}    name           The name of the feed
 * @apiSuccess {String}    description    The description of the feed
 * @apiSuccess {String}    channelId      The channel Id
 * @apiSuccess {String}    applicationId  The Id of the application the feed belongs to
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "name": "Son Goku",
 *       "description": "The Son Goku feed. Here Son Goku will push all his news.",
 *       "channelId": "ASJDH86ASD678ASDASD768",
 *       "applicationId": "KAJFDA786FDS87FDS78F6",
 *       "tagList": [
 *         "Saiyan",
 *         "Dragon Ball",
 *         "Dragon Ball Z"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponseGet() {

}

/**
 * @apiDefine FeedPostResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    name           The name of the feed
 * @apiSuccess {String}    description    The description of the feed
 * @apiSuccess {String}    channelId      The channel Id
 * @apiSuccess {String}    applicationId  The Id of the application the feed belongs to
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "name": "Son Goku",
 *       "description": "The Son Goku feed. Here Son Goku will push all his news.",
 *       "channelId": "ASJDH86ASD678ASDASD768",
 *       "applicationId": "KAJFDA786FDS87FDS78F6",
 *       "tagList": [
 *         "Saiyan",
 *         "Dragon Ball",
 *         "Dragon Ball Z"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePost() {

}

/**
 * @apiDefine FeedPutResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    name           The name of the feed
 * @apiSuccess {String}    description    The description of the feed
 * @apiSuccess {String}    channelId      The channel Id
 * @apiSuccess {String}    applicationId  The Id of the application the feed belongs to
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "name": "Son Goku",
 *       "description": "The Son Goku feed. Here Son Goku will push all his news.",
 *       "channelId": "ASJDH86ASD678ASDASD768",
 *       "applicationId": "KAJFDA786FDS87FDS78F6",
 *       "tagList": [
 *         "Saiyan",
 *         "Dragon Ball",
 *         "Dragon Ball Z"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePut() {

}

/**
 * @apiDefine FeedDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete() {

}
