package org

import (
	"strconv"
	"github.com/feedlabs/elasticfeed/resource"
	"github.com/feedlabs/elasticfeed/service/store/v1/template"
	"errors"
	"sort"
)

func GetEntry(org *resource.Org) (entry map[string]interface{}) {
	entry = make(map[string]interface{})

	stats := make(map[string]string)
	stats["tokens"] = strconv.Itoa(org.Tokens)
	stats["admins"] = strconv.Itoa(org.Admins)
	stats["apps"] = strconv.Itoa(org.Applications)

	entry["id"] = org.Id
	entry["name"] = org.Name
	entry["stats"] = stats

	return entry
}

func GetError(err error) (entry map[string]interface{}, code int) {
	return template.Error(err)
}

func GetSuccess(msg string) (entry map[string]string, code int) {
	return template.Success(msg)
}

type By func(p1, p2 *resource.Org) bool

func (by By) Sort(orgs []*resource.Org) {
	ps := &OrgSorter{
		orgs: orgs,
		by:      by,
	}
	sort.Sort(ps)
}

type OrgSorter struct {
	orgs []*resource.Org
	by      func(p1, p2 *resource.Org) bool
}

func (s *OrgSorter) Len() int {
	return len(s.orgs)
}

func (s *OrgSorter) Swap(i, j int) {
	s.orgs[i], s.orgs[j] = s.orgs[j], s.orgs[i]
}

func (s *OrgSorter) Less(i, j int) bool {
	return s.by(s.orgs[i], s.orgs[j])
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
func ResponseGetList(orgList []*resource.Org, formatter *template.ResponseDefinition) (entryList []map[string]interface{}, code int) {
	var output []map[string]interface{}

	orderby := formatter.GetOrderBy()
	orderdir := formatter.GetOrderDir()
	if orderby == "id" {

		idasc := func(p1, p2 *resource.Org) bool {
			return p1.Id < p2.Id
		}

		iddesc := func(p1, p2 *resource.Org) bool {
			return p1.Id > p2.Id
		}

		if orderdir == "asc" {
			By(idasc).Sort(orgList)
		} else if orderdir == "desc" {
			By(iddesc).Sort(orgList)
		} else {
			errMsg, _ := GetError(errors.New("Unknown ordering direction `" + orderdir + "`"))
			output = append(output, errMsg)
			return output, template.HTTP_CODE_ACCESS_FORBIDDEN
		}

	} else {
		errMsg, _ := GetError(errors.New("Unknown ordering field `" + orderby + "`"))
		output = append(output, errMsg)
		return output, template.HTTP_CODE_ACCESS_FORBIDDEN
	}

	start := formatter.GetPage() * formatter.GetLimit()
	end := start + formatter.GetLimit()
	if len(orgList) < end {
		end = len(orgList)
	}

	if start > end {
		errMsg, _ := GetError(errors.New("Paging is out of range"))
		output = append(output, errMsg)
		return output, template.HTTP_CODE_ACCESS_FORBIDDEN
	}

	orgListPaging := orgList[start:end]

	for _, org := range orgListPaging {
		output = append(output, GetEntry(org))
	}

	return output, template.GetOK()
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
func ResponseGet(org *resource.Org, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(org), template.GetOK()
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
func ResponsePost(org *resource.Org, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(org), template.PostOK()
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
func ResponsePut(org *resource.Org, formatter *template.ResponseDefinition) (entry map[string]interface{}, code int) {
	return GetEntry(org), template.PutOK()
}

/**
 * @apiDefine OrgDeleteResponse
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 */
func ResponseDelete(msg string, formatter *template.ResponseDefinition) (entry map[string]string, code int) {
	return GetSuccess(msg)
}
