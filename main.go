package main

import (
	_ "myproject/routers"
	"github.com/astaxie/beego"
	"myproject/controllers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

