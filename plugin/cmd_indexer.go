package plugin

import (
	"log"
	"github.com/feedlabs/elasticfeed/plugin/model"
)

type cmdIndexer struct {
	indexer model.Indexer
	client  *Client
}

func (b *cmdIndexer) Prepare(config ...interface{}) ([]string, error) {
	defer func() {
		r := recover()
		b.checkExit(r, nil)
	}()

	return b.indexer.Prepare(config...)
}

func (b *cmdIndexer) Run(cache model.Cache) (model.Artifact, error) {
	defer func() {
		r := recover()
		b.checkExit(r, nil)
	}()

	return b.indexer.Run(cache)
}

func (b *cmdIndexer) Cancel() {
	defer func() {
		r := recover()
		b.checkExit(r, nil)
	}()

	b.indexer.Cancel()
}

func (c *cmdIndexer) checkExit(p interface{}, cb func()) {
	if c.client.Exited() && cb != nil {
		cb()
	} else if p != nil && !Killed {
		log.Panic(p)
	}
}
