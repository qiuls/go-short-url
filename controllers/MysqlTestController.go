package controllers

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
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

const PASSWORD string = "******"
const OK string = "ok"
const NO string = "no"

func (c *MysqlTestController) TestJson() {

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

func NewRedisPool() *redis.Client {

	redis_host := beego.AppConfig.String("redis_host")
	redis_port := beego.AppConfig.String("redis_port")
	redis_pass := beego.AppConfig.String("redis_pass")
	redis_db, _ := beego.AppConfig.Int("redis_db")

	add_r := redis_host + ":" + redis_port

	Client := redis.NewClient(&redis.Options{
		Addr:     add_r,
		PoolSize: 1000,
		Password: redis_pass, // no password set
		DB:       redis_db,   // use default DB

		ReadTimeout:  time.Millisecond * time.Duration(100),
		WriteTimeout: time.Millisecond * time.Duration(100),
		IdleTimeout:  time.Second * time.Duration(60),
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic("init redis error")
	} else {
		fmt.Println("init redis ok")
	}
	return Client

}

func (c *MysqlTestController) TestRedis() {

	user_key := c.GetString("user_key")

	if user_key == "" {
		c.Ctx.WriteString("用户参数出错")
		return
	}

	redis := NewRedisPool()
	err := redis.Set(user_key, user_key, 1*time.Second).Err()
	if err != nil {
		fmt.Println("redis set error")
		c.Ctx.WriteString("redis set error")
		return
	}

	val, err := redis.Get(user_key).Result()

	if err != nil {
		fmt.Println("redis get error")
		c.Ctx.WriteString(val)
		return
	}

	fmt.Println(val)
	c.Ctx.WriteString(val)
}
