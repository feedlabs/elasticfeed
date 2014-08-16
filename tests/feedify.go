package main

import (
	"fmt"
	"flag"
	"strconv"
	"github.com/feedlabs/feedify"
	"github.com/astaxie/beego/config"
)

var configPath string

func InitConfig() (config.ConfigContainer) {
	const (
		configFilePath 	= "conf/app.conf"
		usage       	= "set config file path"
	)
	flag.StringVar(&configPath, "config", configFilePath, usage)
	flag.Parse()

	CFPConfig, _ := config.NewConfig("ini", configPath)

	return CFPConfig
}

func main() {
	conf := InitConfig()
	port, _ := strconv.Atoi(conf.String("feedify::port"))

	fmt.Printf("Starting app '%s' on port '%d'\n", conf.String("appname"), port)

	feedify.Run(port)
}
