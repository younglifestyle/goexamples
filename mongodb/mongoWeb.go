package main

import "C"
import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Repository1 struct {
	C *mgo.Collection
}

var coll Repository1

func main() {
	r := gin.Default()

	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	coll = Repository1{C: session.DB("test").C("people")}

	r.POST("/add", addInfo)

	r.Run(":8070")
}

func (r *Repository1) Create(docs ...interface{}) error {
	return r.C.Insert(docs)
}

func (r *Repository1) GetAll() []interface{} {
	var users []interface{}
	iter := r.C.Find(nil).Iter()

	var result interface{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *Repository1) Delete(id string) error {
	return r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

type STest struct {
	IndexNumbering string `json:"index_no" bson:"index_no"`
	testQuery      map[string]interface{}
}

var testData = string(`
{
	"index_no":"123",
	"report_info": {
		"department":"development",
		"reporter":"one",
		"time":"2018-10-17 10:33:00",
		"test1":"1",
		"test3":"1",
		"test2":"1",
		"test212":{
			"one":"123",
			"arew":"321"
		}
	},
	"report_describe": {
		"department":"development",
		"reporter":"one",
		"time":"2018-10-17 10:33:00"
	},
	"test_result": {
		"department":"development",
		"reporter":"one",
		"time":"2018-10-17 10:33:00"
	}
}`)

func addInfo(c *gin.Context) {

	var testQuery map[string]interface{}

	err := c.BindJSON(&testQuery)
	if err != nil {
		log.Println("error :", err, testQuery)
		return
	}

	log.Println("data :", testQuery)

	out, _ := bson.Marshal(testQuery)
	log.Println("out :", string(out))

	err = coll.C.Insert(testQuery)
	log.Println("error1 :", err)

	// 插入json数据
	err = coll.C.Insert(testData)
	log.Println("insert error :", err)

	//coll.C.Insert(out)
	//log.Println("error2 :", err)

	var result map[string]interface{}
	err = coll.C.Find(bson.M{"test": "123"}).One(&result)
	log.Println("error2 :", err, result)

	err = coll.C.Find(bson.M{"name": "Awa"}).
		One(&result)
	log.Println("error2 :", err, result)

	err = coll.C.Update(bson.M{"index_no": "121344",
		"test2.test123.test234": "123"},
		bson.M{"$set": bson.M{"test2.test123.test3.test1": "1191974"}})
	log.Println("err3 :", err)

	var etst map[string]interface{}
	err = coll.C.Find(bson.M{"index_no": "sacsacsa"}).
		One(&etst)
	log.Println("err4 :", err)

	var etst1 []interface{}
	coll.C.Find(nil).Select(bson.M{"_id": 1}).
		Sort("-_id").Limit(5).All(&etst1)
	log.Println("err5 :", etst1)

	err = coll.C.Find(bson.M{"_id": bson.M{"$lt": bson.ObjectIdHex("5bc6dc74c1962da15157f0f9")}}).
		All(&etst1)
	log.Println("error 15 :", etst1, err)

	c.JSON(http.StatusOK, "all is ok")
}
