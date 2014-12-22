package entry


/**
 * @apiDefine EntryGetListByFeedResponse
 *
 * @apiSuccess {Object}    feed
 * @apiSuccess {String}    feed.id                The application Id
 * @apiSuccess {String}    feed.name              The name of the application
 * @apiSuccess {String}    feed.channelId         The channel Id of the feed
 * @apiSuccess {String}    feed.applicationId     The id of the application the feed belongs to
 * @apiSuccess {Int}       feed.createStamp       Unix time stamp of create time
 * @apiSuccess {Object[]}  entryList              Array of all entries
 * @apiSuccess {String}    entryList.id           The entry Id
 * @apiSuccess {Int}       entryList.createStamp  Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "feed": {
 *           "id": "KAJFDA7GFTRE87FDS78F7",
 *           "name": "Son Goku",
 *           "channelId": "ASJDH86ASD678ASDASD768",
 *           "applicationId": "KAJFDA786FDS87FDS78F6",
 *           "createStamp": "1415637736",
 *       }
 *       "entryList": [
 *         {
 *           "id": "KAJFDA7GFTRE87FDS78F7",
 *           "createStamp": "1415637736",
 *         },
 *         ...
 *       ]
 *     }
 */
func ResponseGetListByFeed() {

}

/**
 * @apiDefine EntryGetResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    data           The data of the entry
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
/**
 * @apiDefine EntryGetByFeedResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    feedId         The feed id
 * @apiSuccess {String}    data           The data of the entry
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "feedId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponseGet() {

}

/**
 * @apiDefine EntryPostResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    data           The data of the entry
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePost() {

}

/**
 * @apiDefine EntryPostToFeedResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    feedId         The feed id
 * @apiSuccess {String}    data           The data of the entry
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "feedId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
/**
 * @apiDefine EntryAddToFeedResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    feedId         The feed id
 * @apiSuccess {String}    data           The data of the entry
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "feedId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePostToFeed() {

}

/**
 * @apiDefine EntryPutResponse
 *
 * @apiSuccess {String}    id             The feed Id
 * @apiSuccess {String}    applicationId  The application id
 * @apiSuccess {String}    [feedId]      The feed id
 * @apiSuccess {String}    data           The data of the entry
 * @apiSuccess {String[]}  tagList        List of set tags
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "applicationId": "KAJDFE7GFTRE87FDS78F7",
 *       "feedId": "KAJDFE7GFTRE87FDS78F7",
 *       "data": "Hello, I'm Son Gocu and this is my first post.",
 *       "tagList": [
 *         "First",
 *         "Awesome"
 *       ],
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePut() {

}

/**
 * @apiDefine EntryDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete() {

}

/**
 * @apiDefine EntryRemoveResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseRemove() {

}
