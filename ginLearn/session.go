package main

import (
	"net/http"

	"fmt"
	"log"
	"net/url"

	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store, _ := redis.NewStore(10, "tcp",
		"localhost:6379", "", []byte("secret"))
	store.Options(sessions.Options{
		MaxAge: int(30 * time.Minute),
		Path:   "/",
	})
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}

func test() {
	u := url.Values{}
	u.Set("user_email", "test@mail.com")
	testUrl := "http://" + "127.0.0.1" + "/do?" + u.Encode()
	fmt.Printf("test url is %s\n", testUrl)

	req, err := http.NewRequest("GET", testUrl, nil)
	if err != nil {
		log.Fatalf("generate req of url %s error: %v", testUrl, err)
	}
	req.AddCookie(cookie)
}
