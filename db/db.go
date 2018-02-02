package db

import (
	Conf "apidemo/conf"
	. "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var mysqlConns map[string]*DB = make(map[string]*DB)
var mysqlConnKey = "default"
var mysqlMux sync.Mutex
//func GetConn (key ...string)  {
//
//	if key != nil {
//		dbKey = key[0]
//	}
//
//	conf := conf.GetConf()
//
//	dbMap := conf.Get("db").Get(dbKey).MustMap()
//	switch dbMap["type"] {
//	case "mysql":
//		return getMysqlConn()
//	}
//	gorm.Open()
//}

func GetMysqlConn (key ...string) *DB {
	//mysqlMux.Lock()
	//defer mysqlMux.Unlock()

	if key != nil {
		mysqlConnKey = key[0]
	}

	conf := Conf.GetConf()

	if mysqlConns[mysqlConnKey] == nil {
		mysqlMap := conf.Get("mysql").Get(mysqlConnKey)
		username := mysqlMap.Get("username").MustString()
		password := mysqlMap.Get("password").MustString()
		dbname := mysqlMap.Get("dbname").MustString()
		charset := mysqlMap.Get("charset").MustString()
		params := username + ":" + password + "@/" + dbname + "?charset=" + charset + "&parseTime=True&loc=Local"
		conn, err := Open("mysql", params)
		if err != nil {
			return nil
		}
		mysqlConns[mysqlConnKey] = conn
	}
	return mysqlConns[mysqlConnKey]
}