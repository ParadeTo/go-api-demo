package test

import (
	"testing"
	"apidemo/conf"
)

func Test_Conf1(t *testing.T) {
	conf1 := conf.GetConf("./conf_test.json")
	conf2 := conf.GetConf("./conf_test.json")
	if conf1 != conf2 {
		t.Error("conf1 is not equal to conf2")
	}

	if conf1.Get("port").MustString() != "9090" {
		t.Error("port is not correct")
	}

	if conf1.Get("db").Get("default").Get("username").MustString() != "root" {
		t.Error("db username is not correct")
	}


}

func Test_Conf2(t *testing.T) {
	conf3 := conf.GetConf("./none.json")
	if conf3 != nil {
		t.Error("conf3 should be null")
	}
}
