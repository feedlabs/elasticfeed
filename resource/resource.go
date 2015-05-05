package resource

import (
	"errors"

	"github.com/feedlabs/feedify/service"
	"github.com/feedlabs/feedify/graph"
	"github.com/feedlabs/feedify/stream"
)

const (
	RESOURCE_ORG_LABEL         = "org"
	RESOURCE_ADMIN_LABEL       = "admin"
	RESOURCE_TOKEN_LABEL       = "token"
	RESOURCE_APPLICATION_LABEL = "application"
	RESOURCE_FEED_LABEL        = "feed"
	RESOURCE_ENTRY_LABEL       = "entry"
	RESOURCE_METRIC_LABEL      = "metric"
	RESOURCE_VIEWER_LABEL      = "viewer"
	RESOURCE_WORKFLOW_LABEL    = "workflow"
	RESOURCE_PLUGIN_LABEL      = "plugin"
)

var (
	Orgs              map[string]*Org
	Admins            map[string]*Admin
	Tokens            map[string]*Token
	Applications      map[string]*Application
	Feeds             map[string]*Feed
	Entries           map[string]*Entry
	Metrics           map[string]*Metric
	Viewers           map[string]*Viewer
	Workflows         map[string]*Workflow
	Plugins           map[string]*Plugin

	message    *stream.StreamMessage
	storage    *graph.GraphStorage
)

type Org struct {
	Id               string
	Name             string
	Data             string

	Tokens            int
	Admins            int
	Applications      int
}

type Admin struct {
	Id               string
	Org              *Org

	Username              string
	Maintainer            bool
	Whitelist             []string
	Data                  string

	Tokens                int
}

type Token struct {
	Id                  string
	Admin                *Admin
	Data                string
}

type Application struct {
	Id        string
	Org       *Org
	Data      string
	Feeds     int
}

type Feed struct {
	Id            string
	Application   *Application
	Data          string

	Entries         int
	Workflows       int
}

type Entry struct {
	Id        string
	Feed      *Feed
	Data      string
}

type Viewer struct {}

type Metric struct {}

type Workflow struct {
	Id             string
	Feed           *Feed
	Default        bool
	Data           string
}

type Plugin struct {
	Id            string
	Name          string
	Group         string
	Version       string
	Path          string
	License       string
}

func Contains(s []string, e string) bool {
	for _, a := range s { if a == e { return true } }
	return false
}

func ConvertInterfaceToStringArray(d interface{}) []string {
	data := d.([]interface{})
	output := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		output[i] = data[i].(string)
	}
	return output
}

func InitResources() {
	Admins = make(map[string]*Admin)
	Applications = make(map[string]*Application)
	Feeds = make(map[string]*Feed)
	Entries = make(map[string]*Entry)
	Orgs = make(map[string]*Org)
	Tokens = make(map[string]*Token)
	Metrics = make(map[string]*Metric)
	Viewers = make(map[string]*Viewer)
	Plugins = make(map[string]*Plugin)
}

func InitStorage() {
	graph_service, _ := service.NewGraph()
	if graph_service == nil {
		panic(errors.New("Cannot create graph service"))
	}
	storage = graph_service.Storage
}

func init() {
	InitStorage()
	InitResources()
}
