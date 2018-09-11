package db

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

// "mongodb://localhost:27017/articles_demo_dev"
// "localhost:27017"

// 只针对一个database, 针对多个Db可以参考mysql的实现
func InitMongoDb() {

	dsn := viper.GetString("dsn")
	if dsn == "" {
		log.Fatal("please add dsn in the ./cfgfile/cfg.json")
	}

	mongo, err := mgo.ParseURL(dsn)
	s, err := mgo.Dial(dsn)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", dsn)
	Session = s
	Mongo = mongo
}

// Connect middleware
// Mongo.Database 为"mongodb://localhost:27017/articles_demo_dev"中的articles_demo_dev
// 可以将这个dsn设置成从viper中读取
func Connect(c *gin.Context) {
	s := Session.Copy()

	defer s.Close()

	c.Set("db", s.DB(Mongo.Database))
	c.Next()
}
