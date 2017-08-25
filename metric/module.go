package metric

import (
	//"github.com/tguangch/stats/cncf"
	"github.com/tguangch/gtools/common"
	"errors"
	"github.com/tguangch/gtools/metric/elasticsearch"
	"github.com/tguangch/gtools/metric/redisModule"
	//"fmt"
)

var _mudules = make(map[string]common.MetricModule)

func init(){
	_mudules["elasticsearch"] = elasticsearch.NewModule()
	_mudules["redis"] = redisModule.NewModule()
}

func Registry(moduleNmae string, module common.MetricModule){
	_mudules[moduleNmae] = module
}

func GetModule(module string) (common.MetricModule, error){
	if expect, ok:=_mudules[module]; ok{
		return expect, nil
	}
	return common.MetricModule{}, errors.New("metric module '" + module + "' not found,")
}