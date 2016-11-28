package influxdb

import (
    "testing"
    "fmt"
)

func TestXYZ(t *testing.T) {
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
	err := influxdb.Save(config, data)
	if err != nil {
		fmt.Println("error,", err)
	}
}

