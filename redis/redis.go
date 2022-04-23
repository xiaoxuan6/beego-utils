package redis

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"strconv"
	"sync"
)

var instance cache.Cache
var once = &sync.Once{}

func GetInstance() cache.Cache {

	if instance == nil {

		host := beego.AppConfig.String("host")
		port, _ := beego.AppConfig.Int("port")
		key := beego.AppConfig.DefaultString("key", "default")
		dbNum := beego.AppConfig.DefaultInt("db_num", 0)

		if len(host) < 1 || port < 1 || dbNum < 0 {
			panic("redis 配置信息错误")
		}

		once.Do(func() {
			conn := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))

			rdb, err := cache.NewCache("redis", `{"key":"`+key+`", "conn":"`+conn+`", "dbNum":"`+strconv.Itoa(dbNum)+`"}`)

			if err != nil {
				panic("redis 连接失败")
			}

			instance = rdb
		})
	}

	return instance
}
