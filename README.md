# go-short-url
go bee的缩短url服务

redis 服务
go  -u get github.com/go-redis


host = 'http://host'

# 短url设置url

{host}/short-url
参数
POST

url string  "http://baidu.com"  设置的url
day_num int  100               需要设置的过期天数

响应

{
    "Code": 200,
    "Data": {
        "expire_at": "2021-04-17 13:56:09",
        "key": "l",
        "long_url": "https://blog.csdn.net/iamlihongwei/article/details/79550958",
        "url": "http://127.0.0.1:8082/l"
    },
    "Message": "success"
}

# 短url获取url

{host}/short-url

参数
key string  "d"  响应的短url key

响应
{
    "Code": 200,
    "Data": {
        "expire_at": "2020-12-29 11:31:13",
        "key": "d",
        "long_url": "http://baidu.com",
        "url": "http://127.0.0.1:8082/d"
    },
    "Message": "success"
}

