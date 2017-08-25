package redisModule

import (
	"github.com/tguangch/gtools/common"
	"github.com/tguangch/gtools/cncf"
	"github.com/tguangch/gtools/output/influxdb"
	"strings"
	"github.com/tguangch/gtools/redis"
)

type _metric struct {
	config 		cncf.MetricConfig
	processors	[]common.Processor
	//dataConverter	func(interface{}) interface{}
}

func (this _metric)Start(output common.Output){
	for _, p := range this.processors {
		p.Metric = this
		p.DataConverter = DataConverterToInflux
		p.Output = output
		p.Start()
	}
}

func (this _metric)Collect(host string) (interface{}, error) {
	hostport := strings.Split(host, ":")
	hostOrIp := hostport[0]
	port := hostport[1]
	infoMap, err:= redis.InfoWithMap(hostOrIp, port, "")

	_item := Item{
		host : hostOrIp,
		port : port,
		infoMap : infoMap,
	}
	return _item, err
}

func NewMetric(config cncf.MetricConfig) common.Metric {
	m := _metric{
		config: config,
	}
	_processors := make([]common.Processor, 0)
	for _, host := range config.Hosts {
		_p := common.NewProcessor(host, config.Enabled, config.Period)
		//_p.Collect = m.Collect
		_processors = append(_processors, _p)
	}
	m.processors = _processors
	return m
}

func NewModule() common.MetricModule{
	return common.MetricModule{NewMetric}
}

func  DataConverterToInflux(text interface{}) interface{}{
	item, ok := text.(Item);
	if !ok{
		return nil
	}

	datas := make([]influxdb.Data, 0, 0)

	datas = append(datas,
		influxdb.Data{
			Name : "info",
			Tags : map[string]string{
				"ip" : item.host,
				"port" : item.port,
			},
			Fields : map[string] interface{}{
				"used_memory" : common.AtoInt64(item.infoMap["used_memory"], 0),
				"connected_clients" : common.AtoInt64(item.infoMap["connected_clients"], 0),
				"instantaneous_ops_per_sec" : common.AtoInt64(item.infoMap["instantaneous_ops_per_sec"], 0),
				"instantaneous_input_kbps" : common.AtoFloat64(item.infoMap["instantaneous_input_kbps"],0),
				"instantaneous_output_kbps" : common.AtoFloat64(item.infoMap["instantaneous_output_kbps"], 0),
				"mem_fragmentation_ratio" : common.AtoInt64(item.infoMap["mem_fragmentation_ratio"], 0),
			},
		},
	)
	return datas
}