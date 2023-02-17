package redis

import (
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gomodule/redigo/redis"
)

var rpool *redis.Pool
var once sync.Once

func newPool() *redis.Pool {
	//ctx := context.Background()

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func InitRedis() {
	once.Do(func() {
		rpool = newPool()
	})

}

func Get(key string) (interface{}, error) {
	conn := rpool.Get()
	value, err := redis.String(conn.Do("Get", key))
	if err != nil {
		logs.Info(err)
	}
	return value, err
}

func Set(key string, value interface{}) {
	conn := rpool.Get()
	_, err := redis.String(conn.Do("Set", key, value))
	if err != nil {
		logs.Info(err)
	}
}

func SetNx(key string, value interface{}) {
	conn := rpool.Get()
	_, err := conn.Do("Set", key, value, "NX")
	if err != nil {
		logs.Info(err)
	}
}

func SetEx(key string, value interface{}, time int) {
	conn := rpool.Get()
	_, err := conn.Do("Set", key, value, "EX", time)

	if err != nil {
		logs.Info(err)
	}

}
