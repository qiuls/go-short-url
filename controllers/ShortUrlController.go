package controllers

import (
	"math"
	"myproject/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	// "strconv"
)

type ShortUrlController struct {
	beego.Controller
}

type ShortJsonReturn struct {
	Code    int
	Data    map[string]string
	Message string
}

func (c *ShortUrlController) returnJsonData(code int, message string, data map[string]string) bool {
	if code != 200 {
		data["error"] = "erorr"
	}

	data_json := ShortJsonReturn{Code: code, Data: data, Message: message}
	c.Data["json"] = &data_json
	c.ServeJSON()
	return true
}

/*
  @param  url string
  @param  day_num int
  @retrun   json
*/

func (c *ShortUrlController) Post() {

	data_map := make(map[string]string)

	url_s := c.GetString("url")

	if url_s == "" {

		c.returnJsonData(-1, "url is empty", data_map)
		return
	}

	day_num, err := c.GetInt("day_num")

	if err != nil {
		c.returnJsonData(day_num, "day_num is empty", data_map)
		return
	}
	if day_num <= 1 {
		day_num = 1
	}

	l, _ := time.LoadLocation("Asia/Shanghai")
	// 	fmt.Println(time.Now().In(l))

	now := time.Now().In(l)
	day_hour_num := day_num * 24

	day_hour_string := strconv.Itoa(day_hour_num)

	dd, _ := time.ParseDuration(day_hour_string + "h")
	dd1 := now.Add(dd)

	short_url_model := models.ShortUrls{Url: url_s, Expire_at: dd1, Key: ""}

	id, err := models.AddUrl(&short_url_model)

	if err != nil {
		c.returnJsonData(-1, "add shortUrl erorr", data_map)
		return
	}

	key := transTo62(id)
	short_url_model.Key = key

	num := models.UpdateData(&short_url_model)
	if num == 0 {
		c.returnJsonData(-1, "short_url_model key erorr ", data_map)
		return
	}

	var url = "http://127.0.0.1:8082/" + key

	data_map["url"] = url
	data_map["key"] = key
	data_map["long_url"] = short_url_model.Url
	data_map["expire_at"] = short_url_model.Expire_at.Format("2006-01-02 15:04:05")
	c.returnJsonData(200, "success", data_map)
	return
}

func (c *ShortUrlController) Get() {

	key := c.GetString("key")

	short_url_model := models.ShortUrls{Key: key}

	err := models.FindOneByKey(&short_url_model, "Key")

	data_map := make(map[string]string)

	if err != nil {
		c.returnJsonData(-1, "Key not exiets", data_map)
		return
	}

	var url = "http://127.0.0.1:8082/" + key

	data_map["url"] = url
	data_map["key"] = key
	data_map["long_url"] = short_url_model.Url
	data_map["expire_at"] = short_url_model.Expire_at.Format("2006-01-02 15:04:05")

	c.returnJsonData(200, "success", data_map)
	return
}

//获取北京时间
// func getBeiJjinTime() {
// 	l, _ := time.LoadLocation("Asia/Shanghai")
// 	fmt.Println(time.Now().In(l))

// }

//62进制加密
func transTo62(id int64) string {

	// 1 -- > 1 10-- > a 61-- > Z

	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var shortUrl []byte

	for {

		var result byte

		number := id % 62

		result = charset[number]

		var tmp []byte

		tmp = append(tmp, result)

		shortUrl = append(tmp, shortUrl...)

		id = id / 62
		if id == 0 {

			break
		}

	}
	return string(shortUrl)

}

//62进制解密

func F62ToId(num string, n int64) int64 {

	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var new_num float64

	new_num = 0.0

	nNum := len(strings.Split(num, "")) - 1

	for _, value := range strings.Split(num, "") {

		tmp := float64(strings.Index(charset, string(value)))

		if tmp != -1 {

			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))

			nNum = nNum - 1

		} else {

			break

		}

	}
	return int64(new_num)

}
