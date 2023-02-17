package redis

import (
	"fmt"
	"testing"
)

func TestRedisGet(t *testing.T) {
	InitRedis()
	va := Get("aa")

	fmt.Println(va)
}

func TestRedisSet(t *testing.T) {
	InitRedis()
	Set("vvvv", "uuuuu")

}

func TestRedisSetNx(t *testing.T) {
	InitRedis()
	SetNx("vvvv1", "qqqq")
}
