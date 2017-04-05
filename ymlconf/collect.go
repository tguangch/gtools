package ymlconf

import (
)

type collect struct {
	Elasticsearch		Elasticsearch	`json:"elasticsearch;ommit"`
}

type Elasticsearch struct {
	Host		[]string
	Metricsets	[]string
	Period		int64
}

type Metric struct {
	Host		[]string
	Metricsets	[]string
	Period		int64
}

type Metricset struct {

}
