package conf

import (
	"io/ioutil"
	json "github.com/bitly/go-simplejson"
	"sync"
)

var conf *json.Json
var confPath = "conf.json"
var once sync.Once
var mux sync.Mutex

func GetConf (path ...string) *json.Json {
	//mux.Lock()
	//defer mux.Unlock()

	if len(path) > 0 {
		confPath = path[0]
	}

	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil
	}

	if conf == nil {
		conf = json.New()
		conf.UnmarshalJSON(data)
	}

	//once.Do(func() {
	//
	//})

	return conf
}


