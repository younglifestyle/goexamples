package main

import (
	"encoding/json"
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type (
	User struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name     string        `json:"name"`
		LastName string        `json:"lastname"`
	}

	Person struct {
		Name  string
		Phone string
	}
)

type Repository struct {
	C *mgo.Collection
}

func (r *Repository) Create(docs ...interface{}) error {
	return r.C.Insert(docs)
}

func (r *Repository) GetAll() []interface{} {
	var users []interface{}
	iter := r.C.Find(nil).Iter()

	var result interface{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *Repository) Delete(id string) error {
	return r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

type BindFileIDMgo struct {
	Id                  bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	ProductId           int             `json:"product_id" bson:"product_id"`
	FileIdBeBindProduct []int           `json:"fileid_bind_product" bson:"fileid_bind_product,omitempty"`
	ModuleId            int             `json:"module_id" bson:"module_id"`
	FileIdBeBindModule  []int           `json:"fileid_bind_module" bson:"fileid_bind_module,omitempty"`
	Test                []bson.ObjectId `json:"test" bson:"test,omitempty"`
}

func main() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	r := Repository{
		C: session.DB("test").C("people"),
	}

	//r.C.Insert(&Person{"Heln", "+55 53 8116 9639"},
	//	&Person{"Ana", "+51 53 8116 9639"},
	//	&Person{"A3a", "+51 53 8116 9639"},
	//	&Person{"Awa", "+51 53 8116 9623"})

	//id := bson.NewObjectId()
	//testUser := &User{
	//	Id:       id,
	//	Name:     "qweqwe",
	//	LastName: "unknow",
	//}

	//err = r.C.Insert(testUser)
	//marshal, _ := bson.Marshal(testUser)
	//err = r.C.Insert(bson.M{
	//	"_id":      id,
	//	"name":     "test",
	//	"lastname": "unknow",
	//})
	//fmt.Println("id :", id, err)

	//fmt.Println(r.GetAll())

	//result := Person{}
	var tes map[string]interface{}
	errStr := r.C.Find(bson.M{"index_no": "213214214213213"}).
		Select(bson.M{"_id": 0, "中国.第一": 1}).One(&tes).Error()

	fmt.Println("error gi :", errStr)
	//if err != nil {
	//	panic(err)
	//}

	//err = r.C.Insert(map[string]interface{}{
	//	"_id": bson.NewObjectId(),
	//	"测试":  "1232141",
	//})
	//fmt.Println("error :", err)

	// 查找最后的map
	//data1, exist := tes["中国"]
	//if exist {
	//	fmt.Println(data1)
	//}
	//
	//fmt.Println("result :", result, tes)

	idTest123 := bson.NewObjectId()
	err = r.C.Insert(BindFileIDMgo{
		Id:                 idTest123,
		ProductId:          123,
		FileIdBeBindModule: []int{1, 2, 3},
	})
	fmt.Println("result error :", err)

	idTest := bson.NewObjectId()
	idTest1 := bson.NewObjectId()
	err = r.C.Insert(map[string]interface{}{
		"test": []bson.ObjectId{idTest1},
		"_id":  idTest,
	})
	fmt.Println("result id :", idTest)

	err = r.C.Update(map[string]interface{}{
		"_id": idTest,
	}, map[string]interface{}{
		"$pull": map[string]interface{}{
			"test": idTest1,
		},
	})
	fmt.Println("error :", err)

	pipe := r.C.Pipe([]bson.M{
		{"$match": bson.M{"_id": idTest123}},
		{"$project": bson.M{"count": bson.M{"$size": "$fileid_bind_module"}}},
	})
	resp := bson.M{}
	err = pipe.One(&resp)
	fmt.Println("error is :", err, resp)

	err = r.C.Update(map[string]interface{}{
		"_id": bson.ObjectIdHex("5bebe0cf7f45aa3270c9e532"),
	}, map[string]interface{}{
		"$addToSet": map[string]string{
			"test": "1234567890",
		},
	})
	fmt.Println("error :", err)

	ids := bson.NewObjectId()
	err = r.C.Insert(BindFileIDMgo{
		Test: []bson.ObjectId{ids},
	})

	var bind BindFileIDMgo
	err = r.C.Find(map[string]interface{}{
		"test": map[string]interface{}{
			"$elemMatch": map[string]interface{}{"$eq": ids},
		},
	}).One(&bind)
	fmt.Println("error test :", bind)

	testObjectJson()
}

type TestJson struct {
	TestId bson.ObjectId `json:"test_id"`
}

func testObjectJson() {
	test := TestJson{TestId: bson.NewObjectId()}

	bytes, _ := json.Marshal(test)
	fmt.Println(string(bytes))

	var t1 TestJson
	json.Unmarshal(bytes, &t1)
	fmt.Println(t1)
}
