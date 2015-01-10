package helper

import (
	"github.com/feedlabs/feedify"
)

func GetApiSecret() string {
	return feedify.GetConfigKey("api::secret")
}

func GetAuthType() string {
	return feedify.GetConfigKey("auth::type")
}

func GetAuthRealm() string {
	return feedify.GetConfigKey("auth::realm")
}
