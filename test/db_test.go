package test

import (
	"testing"
	"apidemo/db"
	. "apidemo/controllers"
	"apidemo/conf"
)

func Test_Template(t *testing.T) {
	template := Template{}
	conf.ConfPath = "./conf_test.json"
	conn := db.GetMysqlConn()
	conn.First(&template)
	if template.ID != 1 {
		t.Error("找不到")
	}
}