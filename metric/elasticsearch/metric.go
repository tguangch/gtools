package elasticsearch

import (
	"github.com/tguangch/gtools/common"
	"github.com/tguangch/gtools/cncf"
	"fmt"
	"bufio"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"errors"
	"github.com/tguangch/gtools/output/influxdb"
)

type _metric struct {
	config 		cncf.MetricConfig
	client 		http.Client
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
	url := "http://" + host + "/_nodes/_local/stats"
	resp, err := this.client.Get(url)
	if err != nil {
		return NilCluster, err
	}
	if resp == nil {
		return NilCluster, errors.New("resp nil error")
	}
	if resp.StatusCode < 200 || resp.StatusCode>299 {
		return NilCluster, errors.New(fmt.Sprintf("resp status is invalid, status=%s", resp.StatusCode))
	}

	reader := bufio.NewReader(resp.Body)
	result, _:= ioutil.ReadAll(reader)

	_cluster := Cluster{}
	json.Unmarshal(result, &_cluster)
	_cluster.Host = host

	return _cluster, nil
}

func NewMetric(config cncf.MetricConfig) common.Metric {
	m := _metric{
		config: config,
		client: http.Client{},
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
	clusterInfo, ok := text.(Cluster);
	if !ok{
		return nil
	}

	datas := make([]influxdb.Data, 0, 0)

	for _, node := range clusterInfo.Nodes {
		//mem : heap, non-heap
		mem := node.Jvm.Mem
		datas = append(datas,
			influxdb.Data{
				Name : "mem",
				Tags : map[string] string {
					"cluster" : clusterInfo.ClusterName,
					"host" : clusterInfo.Host,
				},
				Fields : map[string] interface{} {
					"heap_used" : mem.HeapUsedInBytes,
					"heap_used_percent" : mem.HeapUsedPercent,
					"heap_max" : mem.HeapMaxInBytes,
					"no_heap_used" : mem.NoneHeapUsedInBytes,
				},
			},
		)
		//young, survivor, old
		for _name, _p := range node.Jvm.Mem.Pools {
			datas = append(datas,
				influxdb.Data{
					Name : _name,
					Tags : map[string] string {
						"cluster" : clusterInfo.ClusterName,
						"host" : clusterInfo.Host,
					},
					Fields : map[string] interface{} {
						"used" : _p.UsedInBytes,
						"max" : _p.MaxInBytes,
					},
				},
			)
		}


	}

	return datas
}