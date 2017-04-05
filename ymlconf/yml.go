package ymlconf

import (
	"os"
	"io/ioutil"
	"github.com/yaml"
)

type Yml struct {
	Collect		collect
	Output		Output
}

func Parse(ymlFile string) Yml {
	file, err := os.Open(ymlFile)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	yml := Yml{}
	//m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(b, &yml)
	if err != nil {
		panic(err)
	}

	return yml
}