package routers

import (
	"github.com/feedlabs/feedify"
	"github.com/feedlabs/api/public/v1/controllers"
)

func init() {
	feedify.Router("/", &controllers.DefaultController{}, "get:Get")
	feedify.Router("/v1", &controllers.DefaultController{}, "get:Get")
}
