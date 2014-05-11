package controllers

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
	"strings"
	"cfp/src/lib/db"
	"cfp/src/lib/glue"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/config"
)

type ResponseInfo struct {
}

var mongo *db.Mongo
var nsq *glue.Nsq
var neo4j *db.Neo4j
var memcache *db.Memcache

type DefaultController struct {
	beego.Controller
}

func (this *DefaultController) Get() {
	this.Data["json"] = map[string]string{"succes": "ok"}
	this.ServeJson()
}

func GetResponseFormat(input *context.BeegoInput) string {
	format := "json"
	parts := strings.Split(input.Uri(), ".")
	if (len(parts) > 1) {
		format = parts[len(parts) - 1]
	}
	return format
}

func GetRequestParam(input *context.BeegoInput, param string) string {
	return input.Query(param)
}

func init() {

//	var CFPConfig config.ConfigContainer
//	AppPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
//	CFPConfigPath := filepath.Join(AppPath, "conf", "cfp.conf")
//	CFPConfig, _ = config.NewConfig("ini", CFPConfigPath)

//	fmt.Printf("hello   ")
//	fmt.Printf("====%v===", CFPConfig.String("api::port"))

//	port := flag.Int("port", 8080, "http port")
//	flag.Parse()

//	fmt.Printf("====%d===", *port)

	mongo = db.NewMongo()
	mongo.Connect()

	nsq = glue.NewNsq()
	nsq.Connect()

	neo4j = db.NewNeo4j()
	neo4j.Connect()

	memcache = db.NewMemcache()
	memcache.Connect()
}
