package collect

import (
	"github.com/tguangch/gtools/cncf"
	"github.com/tguangch/gtools/common"
	"github.com/tguangch/gtools/metric"
	"fmt"
	"github.com/tguangch/gtools/output"
)

type Collector struct {
	config 		cncf.Yml
	metrics		[]common.Metric
	outputs		[]common.Output
}

var _collectors = make([]Collector, 0)

func Collect(ymls []cncf.Yml){
	for _, yml := range ymls {
		_collectors = append(_collectors, newCollector(yml))
	}

	for _, c := range _collectors {
		for _, m :=range c.metrics{
			if len(c.outputs) > 0 {
				m.Start(c.outputs[0])
			}
		}
	}
}

func newCollector(yml cncf.Yml) Collector{
	c := Collector{config : yml}
	c.metrics = newMetrics(yml.Metric)
	c.outputs = newOutputs(yml.Output)
	return c
}

func newMetrics(mcs []cncf.MetricConfig) []common.Metric {
	ms := make([]common.Metric, 0)
	for _, mc := range mcs {
		module, err := metric.GetModule(mc.Module)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ms = append(ms, module.NewMetric(mc))
	}

	return ms
}

func newOutputs(ocs []cncf.OutputConfig) []common.Output{
	os := make([]common.Output, 0)
	for _, oc := range ocs {
		module, err := output.GetModule(oc.Module)
		if err != nil {
			fmt.Println(err)
			continue
		}
		os = append(os, module.NewOutput(oc))
	}
	return os
}
