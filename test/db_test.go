package test

import (
	"testing"
	"apidemo/db"
	. "apidemo/controllers"
	"fmt"
)

func Test_Template(t *testing.T) {
	template := Template{}
	conn := db.GetMysqlConn()
	conn.First(&template)
	fmt.Println(template)
}