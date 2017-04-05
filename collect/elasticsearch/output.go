package elasticsearch

import (
	//"github.com/tguangch/gtools/ymlconf"
	"github.com/tguangch/gtools/output/influxdb"
//	"strings"
	"fmt"
)

func NewInfluxOutput(config influxdb.Config) InfluxOutput {
	return &influxOutput{
		config : config,
		client : influxdb.NewClient(config),
	}
}

type InfluxOutput interface {
	Output(clusterInfo Cluster)
}

type influxOutput struct {
	config 		influxdb.Config
	client 		*influxdb.InfluxClient
}
func (this *influxOutput) Output(clusterInfo Cluster){
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

	if err:=this.client.BatchSave(datas); err!=nil {
		fmt.Println(err)
	}
}