package influxdb

import (
	"bytes"
	"strconv"
    "time"
    "log"
	"github.com/influxdb/influxdb/client/v2"
)

type Config struct {
	Host	string
	Port	int
	User	string
	PW		string
	Database	string
}

/*
 * http://10.209.16.113:8086
 */
func (config Config) getAddr() string {
	var ab bytes.Buffer
	ab.WriteString("http://")
	ab.WriteString(config.Host)
	ab.WriteString(":")
	ab.WriteString(strconv.Itoa(config.Port))
	
	return ab.String()
}

type Data struct {
	Name	string
	Tags	map[string]string
	Fields	map[string]interface{}
}

func Save(config Config, data Data) error {
    // Make client
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: config.getAddr(),
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    // Create a new point batch
    bp, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  config.Database,
        Precision: "s",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }
	
	pt, err := client.NewPoint(data.Name, data.Tags, data.Fields, time.Now())
    
	if err != nil {
        log.Fatalln("Error: ", err)
    }
	
	bp.AddPoint(pt)

    // Write the batch
    return c.Write(bp)
}

func BatchSave(config Config, datas []Data) error {
    // Make client
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: config.getAddr(),
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    // Create a new point batch
    bp, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  config.Database,
        Precision: "s",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }
	
	now := time.Now()
	
	for _, data := range datas {
		pt, err := client.NewPoint(data.Name, data.Tags, data.Fields, now)
	    
		if err != nil {
	        log.Fatalln("Error: ", err)
	    }
		bp.AddPoint(pt)	
	}

    // Write the batch
    return c.Write(bp)
}
