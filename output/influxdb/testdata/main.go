package main

import (
    "fmt"
    "github.com/tguangch/gtools/output/influxdb"
)

func main() {
	config := influxdb.Config {
		Host : "10.213.12.74",
		Port : 8086,
		Database : "mydb",
	}
	data := influxdb.Data {
		Name : "test",
		Fields : map[string] interface{} {
			"ops" : 6000,
			"mem" : 3000,
			"rx" : 14325,
			"tx" : 32341,
		},
		Tags : map[string] string {
			"host" : "127.0.0.1",
	    	"ip" : "127.0.0.1",
	    	"app" : "12811",
	    	"port" : "12812",
		},
	}
	
	fmt.Println("config", config)
	fmt.Println("data", data)
	//err := influxdb.Save(config, data)
	
	doBatch(config, data)
}

func doSave(config influxdb.Config, data influxdb.Data){
	err := influxdb.Save(config, data)
	if err != nil {
		fmt.Println("error,", err)
	}
}

func doBatch(config influxdb.Config, data influxdb.Data){
	datas := make([]influxdb.Data, 0, 0)
	datas = append(datas, data)
	
	for _, d := range datas {
		fmt.Println(d)
	}
	
	err := influxdb.BatchSave(config, datas)
	if err != nil {
		fmt.Println("error,", err)
	}
}

