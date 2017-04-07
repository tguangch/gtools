package influxdb

import (
	"github.com/tguangch/gtools/cncf"
	"errors"
	"github.com/tguangch/gtools/common"
)

type _output struct {
	outputConfig 	cncf.OutputConfig
	client		*InfluxClient
}

func(this _output) String() string{
	return this.outputConfig.Host
}

func (this _output)Output(text interface{}) error {
	if data, ok := text.(Data); ok {
		this.client.Save(data)
		return nil
	} else if datas, ok := text.([]Data); ok {
		this.client.BatchSave(datas)
		return nil
	}
	return errors.New("invalid output type")
}

func NewOutput(oc cncf.OutputConfig) common.Output{
	op := _output{outputConfig : oc}

	influxConfig := Config{
		Host: oc.Host,
		Port: oc.Port,
		User: oc.User,
		Password: oc.Password,
		Database: oc.Database,
	}
	op.client = NewClient(influxConfig)

	return op
}

func NewModule() common.OutputModule {
	return common.OutputModule{NewOutput}
}