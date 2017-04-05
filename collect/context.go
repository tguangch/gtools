package collect

import (
	"github.com/tguangch/gtools/ymlconf"
	"github.com/tguangch/gtools/collect/elasticsearch"
)

type Context interface {
	Start()
}

type context struct {
	config		ymlconf.Yml
}

func NewContext(yml ymlconf.Yml) Context{
	return &context{config : yml}
}

func (this *context) Start() {
	elasticsearch.NewCollector(this.config.Collect.Elasticsearch, this.config.Output).Collect()
}