package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/test", func(c *gin.Context) {

		bytes, _ := ioutil.ReadAll(c.Request.Body)

		fmt.Println(string(bytes))
	})

	engine.Run(":9110")
}
