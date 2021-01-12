package controllers

import (
	"fmt"
	"myproject/models"
	"time"

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

	c.TplName = "error/404.tpl"
	c.Data["content"] = "很抱歉您访问的地址不存在"

	//判断404的方法是否存在
	request_url := c.Ctx.Request.RequestURI

	key := request_url[1:]

	redis := NewRedisPool()

	val, err := redis.Get(key).Result()

	if err == nil {
		if val == "" {
			return
		}
		c.Ctx.Redirect(302, val)
		return
	}

	//查询是否进行了短url功能
	url := c.FindByKey(key)
	if url != "" {
		fmt.Println("______" + c.Ctx.Request.RequestURI)
		fmt.Println("______" + key)

		url_cache_second, _ := beego.AppConfig.Int64("so_url_cache_second")

		err := redis.Set(key, url, time.Duration(url_cache_second)*time.Second).Err()
		if err != nil {
			fmt.Println("redis set error")
			c.Ctx.WriteString("redis set error")
			return
		}

		c.Ctx.Redirect(302, url)
		return
	}

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

	l, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(l)

	subM := now.Sub(short_url_model.Expire_at)

	diff_num_minutes := subM.Minutes()
	if diff_num_minutes > 1 {
		fmt.Println("当前Url过期 " + short_url_model.Url)
		return ""
	}
	var url = short_url_model.Url
	return url
}
