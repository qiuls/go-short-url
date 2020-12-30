package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
     beego.Router("/mysql", &controllers.MysqlTestController{})

     beego.Router("/short-url", &controllers.ShortUrlController{})

}

