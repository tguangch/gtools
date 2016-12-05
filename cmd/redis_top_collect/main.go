package main

import (
	"os"
	"flag"
	"time"
	"fmt"
	"strconv"
	"log"
	"github.com/tguangch/gtools/machine"
	"github.com/tguangch/gtools/influxdb"
	"github.com/tguangch/gtools/common"	
)

const helpMessage = `Go source code redis_top_collect.
Usage: redis_top_collect <option>

The option argument perform:
	-h	  		influxdb host 
	-p	  		influxdb port
	-d			influxdb dbName
	-n			execute times per second

Example: 
  $ redis_top_collect -h localhost -p 8086 -d mydb -n 10
  
Source code @https://github.com/tguangch/gtools
`

var (
	dbHost = flag.String("h", "localhost", "influxdb host")
	dbPort = flag.Int("p", 8086, "influxdb port")
	dbName = flag.String("d", "mydb", "influxdb dbName")
)

var (
	times = flag.Int("n", 12, "execute times per second")
)

func main(){
	
	flag.CommandLine.Init("collect", flag.ContinueOnError)
	flag.CommandLine.Usage = func(){  }
	err := flag.CommandLine.Parse(os.Args[1:])
	
	if err != nil {
		if err == flag.ErrHelp {
			PrintHelp()
		} 
		os.Exit(1)	
	}
	
	t := *times
	period := 60 / t
	for i := 0; i<t; i++ {
		topJobStart()
		log.Println("task finished, t =", i)
		if i == t-1 {
			break;
		}
		time.Sleep(time.Duration(period) * time.Second)
	}	
}

const NAME = "top"

func topJobStart() {
	
	var config = influxdb.Config{
		Host : *dbHost,
		Port : *dbPort,
		Database : *dbName,
	}
	
	localIp, err := common.Localipv4()
	if err != nil {
		log.Fatalln("error", err)
	}
	
	items := machine.Top()
	
	datas := make([]influxdb.Data, 0, 0)
	
	for _, item := range items {
		datas = append(datas, 
					influxdb.Data{
						Name : NAME,
						Tags : map[string] string {
					    	"ip" : localIp,
					    	"port" : strconv.Itoa(int(item.Port)),
						},
						Fields : map[string] interface{} {
							"cpu" : item.Cpu,
						},
					},
				)
	}
	
	influxdb.BatchSave(config, datas)
}

func PrintUsage(){
	fmt.Println("try 'redis_top_collect -help' for more info")
}

func PrintHelp() {
	fmt.Println(helpMessage)
}