package main

import (
	"flag"
	"github.com/tguangch/gtools/ymlconf"
	"github.com/tguangch/gtools/collect"
)

var (
	configFile = flag.String("c", "/etc/stats/collector.yml", "location of config file")
)

func main(){
	flag.Parse()
	yml := ymlconf.Parse(*configFile)
	collect.NewContext(yml).Start()

	<- make (chan bool)
}
