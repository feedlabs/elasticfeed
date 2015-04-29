package plugin

import (
	"log"
	"github.com/feedlabs/elasticfeed/plugin/model"
)

type cmdPipeline struct {
	pipeline model.Pipeline
	client  *Client
}

func (b *cmdPipeline) Prepare(config ...interface{}) ([]string, error) {
	defer func() {
		r := recover()
		b.checkExit(r, nil)
	}()

	return b.pipeline.Prepare(config...)
}

func (b *cmdPipeline) Run(data interface {}) (interface {}, error) {
	defer func() {
		r := recover()
		b.checkExit(r, nil)
	}()

	return b.pipeline.Run(data)
}

func (b *cmdPipeline) Cancel() {
	defer func() {
		r := recover()
		b.checkExit(r, nil)
	}()

	b.pipeline.Cancel()
}

func (c *cmdPipeline) checkExit(p interface{}, cb func()) {
	if c.client.Exited() && cb != nil {
		cb()
	} else if p != nil && !Killed {
		log.Panic(p)
	}
}
