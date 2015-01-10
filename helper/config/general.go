package config

import (
	"github.com/feedlabs/feedify"
)

func GetApiSuperuser() string {
	return feedify.GetConfigKey("api::su")
}

func GetApiSecret() string {
	return feedify.GetConfigKey("api::key")
}

func GetApiWhitelist() string {
	return feedify.GetConfigKey("api::whitelist")
}

func GetAuthType() string {
	return feedify.GetConfigKey("auth::type")
}

func GetAuthRealm() string {
	return feedify.GetConfigKey("auth::realm")
}
