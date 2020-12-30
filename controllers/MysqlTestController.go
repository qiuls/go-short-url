package controllers

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego"
)

type MysqlTestController struct {
	beego.Controller
}

type JSONS struct {
	//必须的大写开头
	Website string
	Email   string
	Number  int
}

func (c *MysqlTestController) Get() {

	var tmp int = 1
	for i := 0; i < 1; i++ {
		tmp += i
	}
	var data = &JSONS{"beego.me", "astaxie@gmail.com", tmp}

	c.Data["json"] = data

	l, _ := time.LoadLocation("Asia/Shanghai")

	fmt.Println(reflect.TypeOf(l))
	now := time.Now().In(l)
	fmt.Println(reflect.TypeOf(now))

	c.ServeJSON()
}
