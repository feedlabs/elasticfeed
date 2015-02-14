package main

import (
	_ "github.com/feedlabs/elasticfeed/service"
	"github.com/feedlabs/feedify"

	_ "github.com/feedlabs/elasticfeed/elasticfeed/plugin"
)

func main() {
	feedify.SetStaticPath("/static", "public")
	feedify.Run()
}
