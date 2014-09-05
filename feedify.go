package feedify

import (
	// Golang packages
	"fmt"
	"strconv"

	// Beego framework packages
	"github.com/astaxie/beego"

	// feedify packages
	_ "github.com/feedlabs/feedify/lib/api"
	"github.com/feedlabs/feedify/lib/feedify/config"
	_ "github.com/feedlabs/feedify/lib/feedify/stream/adapter/message"
)

func Banner() {
	fmt.Printf("Starting app '%s' on port '%s'\n", config.GetConfigKey("appname"), config.GetConfigKey("feedify::port"))
}

func Run() {
	Banner()

	beego.HttpPort, _ = strconv.Atoi(config.GetConfigKey("feedify::port"))
	beego.Run()
}
