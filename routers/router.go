package routers

import (
	"myproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.MysqlTestController{})

	beego.Router("/short-url", &controllers.ShortUrlController{})

}
