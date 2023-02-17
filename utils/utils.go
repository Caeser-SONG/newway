package utils

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type RedisDB struct {
	Ip   string
	Port string
}
type AppConfig struct {
	Redis RedisDB `toml:"redis"`
}

var (
	once    sync.Once
	appconf AppConfig
	apppath string = "./conf/app.conf"
)

func InitConfs() {

	if _, err := filepath.Abs(apppath); err != nil {
		panic(err)
	}
	once.Do(func() {
		if _, err := toml.DecodeFile(apppath, &appconf); err != nil {
			panic(err)
		}
	})
	fmt.Println(appconf)
}

func GetRedisIp() string {
	return appconf.Redis.Ip
}

func GetRedisPort() string {
	return appconf.Redis.Port
}
