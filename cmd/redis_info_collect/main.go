package main

import (
	"os"
	"bufio"
	"io"
	"flag"
	"strings"
	"fmt"
	"strconv"
	"log"
	"github.com/tguangch/gtools/influxdb"
	"github.com/tguangch/gtools/common"	
	"github.com/tguangch/gtools/redis"
)

const helpMessage = `Go source code redis_top_collect.
Usage: redis_top_collect <option>

The option argument perform:
	-h	  		influxdb host 
	-p	  		influxdb port
	-d			influxdb dbName
	-c			port conf of redis

Example: 
  $ redis_info_collect -h localhost -p 8086 -d mydb -c /etc/port.txt
  
Source code @https://github.com/tguangch/gtools/cmd/redis_info_collect
`

var (
	dbHost = flag.String("host", "localhost", "influxdb host")
	dbPort = flag.Int("port", 8086, "influxdb port")
	dbName = flag.String("db", "mydb", "influxdb dbName")
	configFile = flag.String("conf", "/etc/port.txt", "port conf of redis server")
)
func PrintUsage(){
	fmt.Println("try 'redis_info_collect -help' for more info")
}

func PrintHelp() {
	fmt.Println(helpMessage)
}

var dbConfig influxdb.Config

func main(){
	
	flag.CommandLine.Init("redis_info_collect", flag.ContinueOnError)
	flag.CommandLine.Usage = func(){  }
	err := flag.CommandLine.Parse(os.Args[1:])
	
	if err != nil {
		if err == flag.ErrHelp {
			PrintHelp()
		} 
		os.Exit(1)	
	}
	
//	if len(flag.Args()) == 0 {
//		PrintUsage()
//		os.Exit(1)
//	}
	
	dbConfig = influxdb.Config{
		Host : *dbHost,
		Port : *dbPort,
		Database : *dbName,
	}
	
	fmt.Println(dbConfig)
	
	cFile, err := os.OpenFile(*configFile, os.O_RDONLY, 755)
	if err != nil {
		log.Fatalln(err)
	}
	defer cFile.Close()
	
	reader := bufio.NewReader(cFile)
	
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		sLine := string(line)
		if strings.HasPrefix(sLine, "#") {
			continue
		}
		s := strings.Split(sLine, ":")
		fmt.Println(s)
		if len(s) == 1{  // port
			do("", s[0], "")
		} else if len(s) == 2 { // host(or ip):port
			do(s[0], s[1], "")
		} else if len(s) == 3 { // host(or ip):port:auth
			do(s[0], s[1], s[2])
		}
	}
	
	//go do("10.209.36.193", "10471", "")
	//do("10.209.36.193", "10471", "")
}

func do (hostOrIp string, port string, auth string){
	//redisHost := "w"+port+".wdds.redis.com"
	if port == "" {
		return 
	}
	if hostOrIp == "" {
		master := "w"+port+".wdds.redis.com"
		_do(master, port, auth)
		
		slave := "r"+port+".wdds.redis.com"
		_do(slave, port, auth)
	} else {
		_do(hostOrIp, port, auth)
	}
}

func _do (hostOrIp string, port string, auth string){
	ip, err := common.Remoteipv4(hostOrIp)
	if err != nil {
		log.Println(err)
		return 
		//log.Fatalln(err)
	}
		
	info, _ := redis.InfoWithMap(ip, port, auth)
	used, _ := strconv.Atoi(info["used_memory"])
	
	conf, _ := redis.Config(ip, port, "", "get", "maxmemory")
	log.Println(conf)
	if len(conf) !=2 {
		return
	}
	max, _ := strconv.Atoi(conf[1])
	
	percentage := used*100 / max
	
	data := influxdb.Data{
			Name : "used_mem_percentage",
			Tags : map[string] string {
		    	"ip" : ip,
		    	"port" : port,
			},
			Fields : map[string] interface{} {
				"value" : percentage,
			},
		}
	
	err = influxdb.Save(dbConfig, data)
	if err != nil {
		log.Println(err)
		//log.Fatalln(err)
	} else {
		log.Println("finished")
	}
}

