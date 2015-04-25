package ann

import (
	"github.com/feedlabs/elasticfeed/common"
	"github.com/feedlabs/elasticfeed/plugin"
	"github.com/feedlabs/elasticfeed/plugin/model"
)

type config struct {
	common.ElasticfeedConfig `mapstructure:",squash"`

	tpl *plugin.ConfigTemplate
}

type Pipeline struct {
	config config
}

func (p *Pipeline) Prepare(raws ...interface{}) ([]string, error) {
	return nil, nil
}

func (p *Pipeline) Run(cache model.Cache) (model.Artifact, error) {
	return &Artifact{"aaa", "bbb", 123}, nil
}

func (p *Pipeline) Cancel() {
}
