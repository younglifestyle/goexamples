package main

import (
	"fmt"
	"mime"

	"github.com/go-resty/resty"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	fmt.Println("is test")
	//c.String(http.StatusForbidden,"is test")
}

func getFileData() {
	req := resty.R()

	resp, err := req.Get("http://172.16.9.229:9081/test")
	if err == nil {
		fmt.Println(string(resp.Body()))
		fmt.Println(resp.Header())

		getHeader := resp.Header().Get("Content-Disposition")
		if getHeader != "" {
			mediaType, params, _ := mime.ParseMediaType(getHeader)
			fmt.Println(mediaType)
			fmt.Println(params["filename"])
		}
	}
}

func main() {

	getFileData()
	return

	router := gin.Default()

	//router.GET("/test1", test)

	// need use http.StatusFound
	//router.GET("/red", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/test")
	//})

	router.GET("/test", func(c *gin.Context) {

		c.Header("Content-Disposition", "attachment; filename="+"1.txt")
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Expires", "0")
		// 如果缓存过期了，会再次和原来的服务器确定是否为最新数据，而不是和中间的proxy
		c.Header("Cache-Control", "must-revalidate")
		c.Header("Pragma", "public")

		c.File("./1.txt")

		//c.JSONP(http.StatusOK, gin.H{"hc": "sb"})
		//c.JSON(http.StatusOK, gin.H{"hc": "sb"})
	})

	router.Run(":9081")
}
