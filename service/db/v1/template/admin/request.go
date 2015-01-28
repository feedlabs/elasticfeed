package admin

import (
	"github.com/feedlabs/feedify/context"
)


/**
 * @apiDefine AdminGetListRequest
 *
 */
func RequestGetList(input *context.Input) {

}

/**
 * @apiDefine AdminGetRequest
 *
 * @apiParam {String} adminId  The admin user id
 */
func RequestGet(input *context.Input) {

}

/**
 * @apiDefine AdminPostRequest
 *
 * @apiParam {String}    mail           The E-Mail Address of the admin user
 * @apiParam {Object[]}  roleList       A array of all roles
 * @apiParam {Int}       roleList.id    Role id (see full list at Appendix)
 * @apiParam {String}    roleList.name  Role name
 */
func RequestPost(input *context.Input) {

}

/**
 * @apiDefine AdminPutRequest
 *
 * @apiParam {String}    adminId        The admin user id
 * @apiParam {String}    mail           The E-Mail Address of the admin user
 * @apiParam {Object[]}  roleList       A array of all roles
 * @apiParam {Int}       roleList.id    Role id (see full list at Appendix)
 * @apiParam {String}    roleList.name  Role name
 */
func RequestPut(input *context.Input) {

}

/**
 * @apiDefine AdminDeleteRequest
 *
 * @apiParam {String}  adminId  The admin user id
 */
func RequestDelete(input *context.Input) {

}
