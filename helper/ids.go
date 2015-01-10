package helper

import (
	"github.com/feedlabs/api/resource"
)

func AdminChannelID(admin *resource.Admin) string {
	return GetMd5(admin.Id + admin.Org.Id)
}
