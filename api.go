package main

import (
	_ "github.com/feedlabs/elasticfeed/public/v1"
	"github.com/feedlabs/feedify"
)

func main() {
	feedify.Run()
}
