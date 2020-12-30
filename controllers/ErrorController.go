package controllers

import (
	"fmt"
	"myproject/models"

	"github.com/astaxie/beego"
)

/**
  该控制器处理页面错误请求
*/
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {

	c.Data["content"] = "未经授权，请求要求验证身份"
	c.TplName = "error/404.tpl"
}

func (c *ErrorController) Error403() {
	c.Data["content"] = "服务器拒绝请求"
	c.TplName = "error/404.tpl"
}
func (c *ErrorController) Error404() {

	//判断404的方法是否存在
	request_url := c.Ctx.Request.RequestURI

	key := request_url[1:]

	//查询是否进行了短url功能
	url := c.FindByKey(key)
	if url != "" {
		fmt.Println("______" + c.Ctx.Request.RequestURI)
		fmt.Println("______" + key)
		c.Ctx.Redirect(302, url)
	}
	c.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	c.TplName = "error/404.tpl"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "server error"
	c.TplName = "error/404.tpl"
}
func (c *ErrorController) Error503() {
	c.Data["content"] = "服务器目前无法使用（由于超载或停机维护）"
	c.TplName = "error/404.tpl"
}

func (c *ErrorController) FindByKey(key string) string {
	short_url_model := models.ShortUrls{Key: key}

	err := models.FindOneByKey(&short_url_model, "Key")
	if err != nil {

		return ""
	}

	var url = short_url_model.Url
	return url
}
