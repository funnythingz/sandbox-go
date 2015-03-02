package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/k0kubun/pp"
	"strconv"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	//c.Do("SET", "counter", 0)

	c1, err := redis.String(c.Do("GET", "counter"))
	if err != nil {
		pp.Println("key not found")
	}

	count, _ := strconv.Atoi(c1)
	count = count + 1
	c.Do("SET", "counter", count)
	c2, _ := redis.String(c.Do("GET", "counter"))
	pp.Println(c2)
}
