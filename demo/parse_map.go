package main
import (
	"fmt"
	"encoding/json"
)

var jsonstr = `
{
  "port": "8080",
  "mysql": {
    "default": {
      "username": "root",
      "password": "123456",
      "dbname": "operation",
      "charset": "utf8"
    },
	"test": {
		"username": "root",
      "password": "123456",
      "dbname": "operation",
      "charset": "utf8"
	}
  }
}
`

type Conf struct {
	Port string `json:"port"`
	Mysql map[string](map[string]string) `json:"mysql"`
}

func main() {
	c := Conf{}
	json.Unmarshal([]byte(jsonstr), &c)
	fmt.Println(c.Mysql["test"]["username"])
}