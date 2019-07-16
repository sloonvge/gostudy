package main

import (
	"fmt"

	redis "github.com/gomodule/redigo/redis"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost: 6379")
		},
	}

	c := pool.Get()
	_, err := c.Do("set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("get data:", r)

}
