package common

import (
	"github.com/feedlabs/elasticfeed/resource"
)

func AdminChannelID(admin *resource.Admin) string {
	return GetMd5(admin.Id + admin.Org.Id)
}
