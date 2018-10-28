package main

import (
	"fmt"

	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error", err)
		return
	}
	defer conn.Close()

	//normalWR(conn)
	//setExpiredTime(conn)
	deleteKey(conn)
}

func normalWR(conn redis.Conn) {
	_, err := conn.Do("SET", "mykey", "super")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return
	}

	username, err := redis.String(conn.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed", err)
	} else {
		fmt.Println("get mykey :", username)
	}
}

func setExpiredTime(conn redis.Conn) {
	_, err := conn.Do("SET", "mykey", "super", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return
	}

	username, err := redis.String(conn.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed", err)
	} else {
		fmt.Println("get mykey :", username)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(conn.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed", err)
	} else {
		fmt.Println("get mykey :", username)
	}
}

func deleteKey(conn redis.Conn) {
	normalWR(conn)

	_, err := conn.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delete failed", err)
		return
	}
	username, err := redis.String(conn.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed", err)
	} else {
		fmt.Println("get mykey :", username)
	}
}
