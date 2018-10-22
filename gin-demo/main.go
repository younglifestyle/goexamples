package main

import (
	"flag"
	"fmt"
	"goexamples/gin-demo/config"
	"goexamples/gin-demo/dbops/db"
	"goexamples/gin-demo/handlers/articles"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 使用MongoDB
func init() {
	db.InitMongoDb()
}

func RegisterHandlers() *gin.Engine {
	log.Printf("preparing to run the server\n")

	router := gin.Default()

	router.Static("/image", "./static/image")
	router.NoRoute(page404)
	router.Use(db.Connect)

	// Articles
	router.GET("/new", articles.New)
	router.GET("/articles/:_id", articles.Edit)
	router.GET("/articles", articles.List)
	router.POST("/articles", articles.Create)
	router.POST("/articles/:_id", articles.Update)
	router.POST("/delete/articles/:_id", articles.Delete)

	return router
}

func main() {

	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	if *version {
		fmt.Println(config.VERSION)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// set log level
	config.InitLog(viper.GetString("log_level"))

	httpAddr := viper.GetString("http.listen")
	if !viper.GetBool("http.enabled") || httpAddr == "" {
		return
	}

	router := RegisterHandlers()
	router.Run(httpAddr)
}
