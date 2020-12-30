package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//由于model这个名字叫 ShortUrls 那么操作的表其实 short_urls
type ShortUrls struct {
	Id int64

	Url string

	Created_at time.Time `orm:"auto_now_add;type(datetime)"`

	Updated_at time.Time `orm:"auto_now;type(datetime)"`

	Expire_at time.Time `orm:"type(datetime)"`

	Key string
}

func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/json?charset=utf8&parseTime=True&loc=Local", 30)
	orm.RegisterModel(new(ShortUrls))
}

func Test(shor_urls *ShortUrls) {
	fmt.Println(shor_urls)
	fmt.Println("---------------------------------")
}

func AddUrl(shor_urls *ShortUrls) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(shor_urls)
	return id, err
}

func UpdateData(shor_urls *ShortUrls) int64 {
	o := orm.NewOrm()

	var Zero int64 = 0


	num, err := o.Update(shor_urls)
	if err == nil {
		return num
	}
	return Zero
}

func FindOne(shor_urls *ShortUrls) error {
	o := orm.NewOrm()
	err := o.Read(shor_urls)
	return err
}

func FindOneByKey(shor_urls *ShortUrls, field string) error {

	o := orm.NewOrm()
	err := o.Read(shor_urls, field)
	return err
}
