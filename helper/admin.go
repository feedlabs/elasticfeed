package helper

import (
	"github.com/feedlabs/api/resource"
)

func IsSuperUser(admin *resource.Admin) bool {
	return admin.Username == GetApiSuperuser()
}

func GetAdminWhitelist(admin *resource.Admin) []string {
	return admin.Whitelist
}

func GetAdminByName(username string) *resource.Admin {
	org := &resource.Org{"0", "", "", 0, 0, 0}
	whitelist := []string{"127.0.0.1", "192.168.1.51"}

	return &resource.Admin{"0", org, username, whitelist, "hello", 0}
}
