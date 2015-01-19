package main

import (
	_ "github.com/feedlabs/elasticfeed/public/v1"
	_ "github.com/feedlabs/elasticfeed/stream"
	"github.com/feedlabs/feedify"
)

func main() {
	feedify.Run()
}
