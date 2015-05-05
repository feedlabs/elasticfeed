package model

import (
	"github.com/feedlabs/elasticfeed/service/stream"
)

type ServiceManager interface {

	GetStreamService() *stream.StreamService

	Init()
}
