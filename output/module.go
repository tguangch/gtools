package output

import (
	"github.com/tguangch/gtools/common"
	"github.com/tguangch/gtools/output/influxdb"
	"github.com/tguangch/gtools/output/elasticsearch"
	"errors"
)

const Module_ElasticSearch = "elasticsearch"
const Module_InfluxDB = "influxdb"
const Module_OpenTsDB = "optsdb"

var _outputs = make(map[string]common.OutputModule)

func init(){
	_outputs[Module_ElasticSearch] = elasticsearch.NewModule()
	_outputs[Module_InfluxDB] = influxdb.NewModule()
	//_outputs[Module_OpenTsDB] = opentsdb.DefaultDatabase
}

func GetModule(module string) (common.OutputModule, error){
	if expect, ok:=_outputs[module]; ok{
		return expect, nil
	}
	return common.OutputModule{}, errors.New("not found")
}