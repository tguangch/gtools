package ymlconf

import (
	"github.com/tguangch/gtools/output/influxdb"
)
type Output struct {
	Influxdb	influxdb.Config		`json:"influxdb"`
}
//
//type Influxdb struct {
//	influxdb.Config
//}
