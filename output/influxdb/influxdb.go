package influxdb

import (
	"bytes"
	"strconv"
    	"time"
    	"log"
	"github.com/influxdb/influxdb/client/v2"
	//"github.com/astaxie/beego/config"
)

type Config struct {
	Host	string			`json:"host"`
	Port	int			`json:"port"`
	User	string			`json:"user"`
	PW		string		`json:"pw"`
	Database	string		`json:"database"`
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

type InfluxClient struct {
	config 		Config
	httpClient 	client.Client
}

func NewClient(config Config) *InfluxClient{
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: config.getAddr(),
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	return &InfluxClient{
		config : config,
		httpClient : c,
	}
}

func (this *InfluxClient) Save(data Data) error {
    c := this.httpClient

    // Create a new point batch
    bp, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  this.config.Database,
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

func (this *InfluxClient) BatchSave(datas []Data) error {
	c := this.httpClient
    // Create a new point batch
    bp, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  this.config.Database,
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
