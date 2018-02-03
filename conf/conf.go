package conf

import (
	"io/ioutil"
	json "github.com/bitly/go-simplejson"
	"sync"
	"fmt"
)

var conf *json.Json
var ConfPath = "conf.json"
var once sync.Once
var mux sync.Mutex

func init() {
	fmt.Println("i am conf")
}

func GetConf (path ...string) *json.Json {
	mux.Lock()
	defer mux.Unlock()

	if len(path) > 0 {
		ConfPath = path[0]
	}

	data, err := ioutil.ReadFile(ConfPath)
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


