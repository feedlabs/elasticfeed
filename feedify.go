package feedify

import (
	"github.com/astaxie/beego"
	_ "github.com/feedlabs/feedify/api"
)

func Run(port int) {
	beego.HttpPort = port
	beego.Run()
}
