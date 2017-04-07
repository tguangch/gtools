package common

import (
	"time"
	"fmt"
)

type Processor struct {
	host 		string
	enabled		bool
	period		int64
	Metric 		Metric
	//Collect 	func(host string) (interface{}, error)
	//Output 		func(interface{}) error
	Output		Output
	DataConverter	func(interface{}) interface{}
}

func (this Processor) Start() {
	go this.Process()
}

func (this Processor) Host() string {
	return this.host
}

func (this Processor) Process(){
	if !this.enabled {
		return
	}

	p := this.period
	if p < 1 {
		p = 1
	}
	d := time.Second * time.Duration(p)

	for {
		select {
		case <- time.After(d):
			stats, err := this.Metric.Collect(this.host)
			err2 := this.Output.Output(this.DataConverter(stats))
			fmt.Println(time.Now(), " ||| collect source:", this.host, ",err:", err," ||| output:", this.Output, " err:", err2)
		}
	}
}

func NewProcessor(host string, enabled bool, period int64) Processor{
	return Processor{host:host, enabled:enabled, period:period}
}

