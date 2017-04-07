package main

import (
	"flag"
	"github.com/tguangch/gtools/cncf"
	_ "github.com/tguangch/gtools/output"
	_ "github.com/tguangch/gtools/metric"
	"github.com/tguangch/gtools/collect"
	"fmt"
	"encoding/json"
)

var (
	configFile = flag.String("c", "/etc/stats/collect.yml", "location of config file")
)

func main(){
	flag.Parse()
	ymls := cncf.Parse(*configFile)

	_j, _ := json.Marshal(ymls)
	fmt.Println(string(_j))

	collect.Collect(ymls)
	<- make (chan bool)
}


