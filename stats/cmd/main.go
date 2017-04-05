package main

import (
	"strconv"
	"fmt"
	"os"
	"log"
	"time"
	"flag"
	"github.com/tguangch/gtools/machine"
	"github.com/tguangch/gtools/influxdb"
	"github.com/tguangch/gtools/common"
)

const useHelp = "Run 'stats -help' for more information.\n"

const helpMessage = `Go source code stats.`

var (
	dbHost = flag.String("h", "localhost", "influxdb host")
	dbPort = flag.Int("p", 8086, "influxdb port")
	dbName = flag.String("db", "mydb", "influxdb dbName")
	configFile = flag.String("c", "configFile", "conf file")
	period = flag.Int("n", 10, "period(unit:s)")
)

func main() {
	
	flag.Usage = func() { fmt.Fprint(os.Stderr, useHelp) }
	flag.CommandLine.Init(os.Args[0], flag.ContinueOnError) // hack
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		// (err has already been printed)
		if err == flag.ErrHelp {
			printHelp()
		}
		os.Exit(2)
	}
	
	p := *period
	
	times := 60 / p
	
	for i := 0; i<times; i++ {
		topJobStart()
		log.Println("task finished, t =", i)
		if i == times-1 {
			break;
		}
		time.Sleep(time.Duration(p) * time.Second)
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

func printHelp() {
	fmt.Fprintln(os.Stderr, helpMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
}