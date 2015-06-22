package config

import (
	"fmt"
	"os"
	"strconv"
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

func GetPluginStoragePath() string {
	return feedify.GetConfigKey("plugin-manager::storage")
}

func GetPluginPortMin() uint {
	port, err := strconv.ParseUint(feedify.GetConfigKey("plugin-manager::port_min"), 10, 0)

	if err == nil {
		return uint(port)
	}

	return 40000
}

func GetPluginPortMax() uint {
	port, err := strconv.ParseUint(feedify.GetConfigKey("plugin-manager::port_max"), 10, 0)

	if err == nil {
		return uint(port)
	}

	return 41000
}

func GetHomeAbsolutePath() string {
	pwd, _ := os.Getwd()
	return pwd
}

func init() {

	_, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot read working directory path!")
		os.Exit(1)
	}

	path := GetHomeAbsolutePath() + "/" + GetPluginStoragePath()
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0777)
			if err != nil {
				fmt.Println("Cannot create plugins storage directory!")
				os.Exit(1)
			}
		}
	}

}
