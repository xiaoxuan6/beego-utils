package redis

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"sync"
)

var instance cache.Cache
var once = &sync.Once{}

func GetInstance() cache.Cache {

	if instance == nil {

		once.Do(func() {
			rdb, err := cache.NewCache("redis", `{"key":"default", "conn":"127.0.0.1:6379", "dbNum":"0"}`)

			if err != nil {
				logs.GetLogger().Println("redis 连接失败", err.Error())
			}

			instance = rdb
		})
	}

	return instance
}
