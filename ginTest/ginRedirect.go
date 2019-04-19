package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	fmt.Println("is test")
	//c.String(http.StatusForbidden,"is test")
}

func main() {
	router := gin.Default()

	router.GET("/test1", test)

	// need use http.StatusFound
	//router.GET("/red", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/test")
	//})

	router.GET("/test", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{"hc": "sb"})
		//c.JSON(http.StatusOK, gin.H{"hc": "sb"})
	})

	router.Run(":8081")
}
