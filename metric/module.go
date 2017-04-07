package metric

import (
	//"github.com/tguangch/stats/cncf"
	"github.com/tguangch/gtools/common"
	"errors"
	"github.com/tguangch/gtools/metric/elasticsearch"
	//"fmt"
)

const Module_ElasticSearch = "elasticsearch"

var _mudules = make(map[string]common.MetricModule)

func init(){
	_mudules[Module_ElasticSearch] = elasticsearch.NewModule()
}

func GetModule(module string) (common.MetricModule, error){
	if expect, ok:=_mudules[module]; ok{
		return expect, nil
	}
	return common.MetricModule{}, errors.New("metric module '" + module + "' not found,")
}