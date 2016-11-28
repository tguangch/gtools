package main

import (
	"strconv"
	"fmt"
	"os"
	"log"
	"flag"
	"github.com/tguangch/gtools/stimer"
	"github.com/tguangch/gtools/machine"
	"github.com/tguangch/gtools/influxdb"
)

const useHelp = "Run 'stats -help' for more information.\n"

const helpMessage = `Go source code stats.`

var (
	dbHost = flag.String("h", "localhost", "influxdb host")
	dbPort = flag.String("p", "8086", "influxdb port")
	dbName = flag.String("db", "mydb", "influxdb dbName")
	configFile = flag.String("c", "configFile", "conf file")
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
	
	//job-1 : 采集redis top 信息，并存入influxdb
	stimer.SimpleTimer{
		Period : 5,
		Task : topJobStart,
	}.ExecTask()
	
	//job-2 : 

}

const NAME = "top"

func topJobStart() {
	
	iPort, err := strconv.Atoi(*dbPort)
	if err != nil {
		log.Fatalln("invalid port")
		os.Exit(2)
	}
	
	var config = influxdb.Config{
		Host : *dbHost,
		Port : iPort,
		Database : *dbName,
	}
	
	localIp, err := Localipv4()
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