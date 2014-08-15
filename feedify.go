package feedify

import (
	"flag"
	"github.com/astaxie/beego"
	_ "github.com/feedlabs/feedify/api"
)

func init() {
	const (
		defaultPort = 9090
		usage       = "set port"
	)
	flag.IntVar(&beego.HttpPort, "port", defaultPort, usage)
}

func Run(port int) {
	flag.Parse()
	beego.HttpPort = port
	beego.Run()
}
