package cncf

import (
	"os"
	"io/ioutil"
	"github.com/yaml"
	"errors"
	"fmt"
)

type Yml struct {
	Metric			[]MetricConfig
	Output			[]OutputConfig
	Include			[]string
}

type MetricConfig struct {
	Module		string		//elasticsearch, redis, ...
	Hosts		[]string	//pairs of ip:port
	Metricset	[]string
	Enabled		bool
	Outputs		[]string
	Period		int64
}

type OutputConfig struct {
	Module		string
	Host		string
	Port		int
	Database	string
	Index		string
	User		string
	Password	string
}

func Parse(ymlFile string) []Yml {
	yml, err := parseBase(ymlFile)
	if err != nil {
		panic(err)
	}

	r := append(make([]Yml, 0), yml)

	// handle include config
	if len(yml.Include) > 0 {
		r = parseInclude(r, yml.Include)
	}

	return r
}

//prevent cycling include
var _allymls = make(map[string]Yml)

func parseBase(ymlFile string) (Yml, error) {
	yml := Yml{}

	file, err := os.Open(ymlFile)
	if err != nil {
		return yml, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return yml, err
	}

	err = yaml.Unmarshal(b, &yml)
	if err != nil {
		return yml, err
	}

	//check if config file already loaded
	if _, ifexist:=_allymls[ymlFile]; ifexist{
		return yml, errors.New(ymlFile + ", already exist!!! check if cycle include")
	}
	_allymls[ymlFile] = yml

	return yml, nil
}

func parseInclude(r []Yml, includeFiles []string) []Yml{
	for _, f := range includeFiles{
		yml, err := parseBase(f)
		if err != nil {
			fmt.Println(err)
			continue
		}
		r = append(r, yml)
		// handle include config
		if len(yml.Include) > 0 {
			r = parseInclude(r, yml.Include)
		}
	}

	return r
}