package elasticsearch

import (
	"github.com/tguangch/gtools/ymlconf"
	"time"
	"net/http"
	"fmt"
)

func NewCollector(c ymlconf.Elasticsearch, op ymlconf.Output) *Collector{

	return &Collector{
		config : c,
		output : op,
		items : make(map[string] CollectorItem, 0),
	}
}

type Collector struct {
	config 		ymlconf.Elasticsearch
	output 		ymlconf.Output
	items		map[string] CollectorItem
	duration 	time.Duration
}

func (this *Collector) Collect(){
	_config := this.config

	period := _config.Period
	if period <= 0 {
		period = 1   // 1 second
	}
	_duration := time.Second * time.Duration(period)

	for _, host := range _config.Host {

		if _, ifExist := this.items[host]; ifExist {
			fmt.Println(host, "already exist!!!")
			continue
		}

		_item := CollectorItem{
			client : http.Client{},
			host : host,
			output : NewInfluxOutput(this.output.Influxdb),
			duration : _duration,
		}

		this.items[host] = _item
		go _item.Collect()
	}
}



