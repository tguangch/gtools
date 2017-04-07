package main

import (
	"os"
	"flag"
	//"time"
	"fmt"
	"github.com/tguangch/gtools/cmd/redis_top_collectd/run"
)

const helpMessage = `Go source code redis_top_collectd.
Usage: redis_top_collectd <mode> <option>

The mode argument determines perform:
	start	  	start the program
	stop	  	stop the program
	status	  	show the status

The option argument perform:
	-c	  		location of conf file 
	-o	  		location of output

Example: 
  $ redis_top_collectd -c /etc/stats/collect.yml -o /var/run/collect.log
`

var (
	pidfile	= flag.String("pidfile", "/var/run/stats/collector.pid", "Write process ID to a file")
	configFile = flag.String("c", "/etc/stats/collector.yml", "location of config file")
	output = flag.String("o", "/var/run/collector.log", "location of output")
)

func main(){
	
	flag.CommandLine.Init("test", flag.ContinueOnError)
	flag.CommandLine.Usage = func(){  }
	err := flag.CommandLine.Parse(os.Args[1:])
	
	if err != nil {
		if err == flag.ErrHelp {
			PrintHelp()
		} else {
			os.Exit(1)
		}		
	}
	
	if len(flag.Args()) == 0 {
		app.Start()
	} else if len(flag.Args()) > 0 {
		switch flag.Arg(0) {
			case "start" :
				app.Start()
			case "stop" :
				app.Stop()
			case "status" :
				app.Status()				
			default :
				fmt.Println("Invalid param", flag.Args())
				PrintUsage()
				os.Exit(1)
		}		
	}

}

func PrintUsage(){
	fmt.Println("'redis_top_collectd -help' for more info")
}

func PrintHelp() {
	fmt.Println(helpMessage)
}