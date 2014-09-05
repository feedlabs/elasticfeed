package config

import (
	"flag"
	"github.com/astaxie/beego/config"
)

var (
	configPath string
	httpPort   int
	conf       config.ConfigContainer
)

func InitConfig() {
	if conf == nil {
		const (
			defaultConfigPath = "conf/app.conf"
			defaultHttpPort   = 8080
		)
		flag.StringVar(&configPath, "config", defaultConfigPath, "set config file path")
		flag.IntVar(&httpPort, "port", defaultHttpPort, "set http port")
		flag.Parse()

		conf, _ = config.NewConfig("ini", configPath)
	}
}

func GetConfigKey(key string) string {
	return conf.String(key)
}

func init() {
	InitConfig()
}
