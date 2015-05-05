package main

import (
	"github.com/feedlabs/elasticfeed/elasticfeed"
)

func main() {
	engine := elasticfeed.NewElasticfeed()
	engine.Run()
}
