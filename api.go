package main

import (
	_ "github.com/feedlabs/elasticfeed/service"
	"github.com/feedlabs/feedify"
)

func main() {
	feedify.Run()
}
