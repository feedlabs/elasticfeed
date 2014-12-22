package admin

/**
 * @apiDefine AdminGetListResponse
 *
 * @apiSuccess {Object[]}  adminList                Array of all admin users
 * @apiSuccess {String}    adminList.id             The id
 * @apiSuccess {String}    adminList.mail           The E-Mail Address of the admin user
 * @apiSuccess {Object[]}  adminList.roleList       A array of all roles
 * @apiSuccess {Int}       adminList.roleList.id    Role id (see full list at Appendix)
 * @apiSuccess {String}    adminList.roleList.name  Role name
 * @apiSuccess {Int}       adminList.createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     "adminList": [
 *       {
 *         "id": "KAJFDA7GFTRE87FDS78F7",
 *         "mail": "son.goku@dragonball.com",
 *         "roleList": [
 *           {
 *             "id": "2",
 *             "name": "admin"
 *           },
 *           ...
 *         "createStamp": "1415637736",
 *       },
 *       ...
 *     ]
 *   }
 */
func ResponseGetList() {

}

/**
 * @apiDefine AdminGetResponse
 *
 * @apiSuccess {String}    id             The id
 * @apiSuccess {String}    mail           The E-Mail Address of the admin user
 * @apiSuccess {Object[]}  roleList       A array of all roles
 * @apiSuccess {Int}       roleList.id    Role id (see full list at Appendix)
 * @apiSuccess {String}    roleList.name  Role name
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "mail": "son.goku@dragonball.com",
 *       "roleList": [
 *         {
 *           "id": "2",
 *           "name": "admin"
 *         },
 *         ...
 *        "createStamp": "1415637736",
 *     }
 */
func ResponseGet() {

}

/**
 * @apiDefine AdminPostResponse
 *
 * @apiSuccess {String}    id             The id
 * @apiSuccess {String}    mail           The E-Mail Address of the admin user
 * @apiSuccess {Object[]}  roleList       A array of all roles
 * @apiSuccess {Int}       roleList.id    Role id (see full list at Appendix)
 * @apiSuccess {String}    roleList.name  Role name
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "mail": "son.goku@dragonball.com",
 *       "roleList": [
 *         {
 *           "id": "2",
 *           "name": "admin"
 *         },
 *         ...
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePost() {

}

/**
 * @apiDefine AdminPutResponse
 *
 * @apiSuccess {String}    id             The id
 * @apiSuccess {String}    mail           The E-Mail Address of the admin user
 * @apiSuccess {Object[]}  roleList       A array of all roles
 * @apiSuccess {Int}       roleList.id    Role id (see full list at Appendix)
 * @apiSuccess {String}    roleList.name  Role name
 * @apiSuccess {Int}       createStamp    Unix time stamp of create time
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 *     {
 *       "id": "KAJFDA7GFTRE87FDS78F7",
 *       "mail": "son.goku@dragonball.com",
 *       "roleList": [
 *         {
 *           "id": "2",
 *           "name": "admin"
 *         },
 *         ...
 *       "createStamp": "1415637736",
 *     }
 */
func ResponsePut() {

}

/**
 * @apiDefine AdminDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete() {

}
