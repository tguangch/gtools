package elasticsearch

import (
	"time"
	"fmt"
	"bufio"
	"io/ioutil"
	"encoding/json"
	"errors"
	"net/http"
)

type CollectorItem struct {
	host		string
	client		http.Client
	output 		InfluxOutput
	duration 	time.Duration
}

func (this *CollectorItem) Collect(){
	for {
		select {
		case <- time.After(this.duration):
			_c, err := this.doCollect()
			fmt.Println(time.Now(), this.host, err)
			if err != nil {
				continue
			}
			this.output.Output(_c)
		}
	}
}

//http://10.213.54.76:10200/_nodes/_local/stats/
func (this *CollectorItem) doCollect () (Cluster, error){
	url := "http://" + this.host + "/_nodes/_local/stats"
	resp, err := this.client.Get(url)
	if err != nil {
		return NilCluster, err
	}
	if resp == nil {
		return NilCluster, errors.New("resp nil error")
	}
	if resp.StatusCode < 200 || resp.StatusCode>299 {
		return NilCluster, errors.New(fmt.Sprintf("resp status is invalid, status=%s", resp.StatusCode))
	}

	reader := bufio.NewReader(resp.Body)
	result, _:= ioutil.ReadAll(reader)

	_cluster := Cluster{}
	json.Unmarshal(result, &_cluster)
	_cluster.Host = this.host

	return _cluster, nil
}