package main

import (
	"fmt"
	"flag"
	"strconv"
	"github.com/feedlabs/feedify"
	"github.com/astaxie/beego/config"
)

var configPath string
var httpPort int

func InitConfig() (config.ConfigContainer) {
	const (
		defaultConfigPath 	= "conf/app.conf"
		defaultHttpPort		= 8080
	)
	flag.StringVar(&configPath, "config", defaultConfigPath, "set config file path")
	flag.IntVar(&httpPort, "port", defaultHttpPort, "set http port")
	flag.Parse()

	CFPConfig, _ := config.NewConfig("ini", configPath)

	return CFPConfig
}

func main() {
	conf := InitConfig()
	port, err := strconv.Atoi(conf.String("feedify::port"))

	if (err == nil) {
		httpPort = port
	}

	fmt.Printf("Starting app '%s' on port '%d'\n", conf.String("appname"), httpPort)

	feedify.Run(httpPort)
}
