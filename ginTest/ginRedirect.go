package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func test(c *gin.Context) {
	fmt.Println("is test")
	//c.String(http.StatusForbidden,"is test")
}

func main() {
	router := gin.Default()

	router.GET("/test", test)

	// need use http.StatusFound
	router.GET("/red", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/test")
	})

	router.Run(":8081")
}