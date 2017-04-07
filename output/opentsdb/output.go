package opentsdb

import (
	"github.com/tguangch/gtools/cncf"
	"github.com/tguangch/gtools/common"
)

type _output struct {

}

func NewOutput(oc cncf.OutputConfig) common.Output{

	return nil
}

func NewModule() common.OutputModule {
	return common.OutputModule{NewOutput}
}