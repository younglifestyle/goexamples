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

type TwoParam struct {
	Test1 int           `json:"test1" bson:"test1"`
	Test2 bson.ObjectId `json:"test2" bson:"test2"`
}

type BindFileIDMgo struct {
	Id                  bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	ProductId           int             `json:"product_id" bson:"product_id"`
	FileIdBeBindProduct []int           `json:"fileid_bind_product" bson:"fileid_bind_product,omitempty"`
	ModuleId            int             `json:"module_id" bson:"module_id"`
	FileIdBeBindModule  []int           `json:"fileid_bind_module" bson:"fileid_bind_module,omitempty"`
	Test                []bson.ObjectId `json:"test" bson:"test,omitempty"`
	ArrTest             []TwoParam      `json:"testArr" bson:"testArr,omitempty"`
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

	//idTest123 := bson.NewObjectId()
	//err = r.C.Insert(BindFileIDMgo{
	//	Id:                 idTest123,
	//	ProductId:          123,
	//	FileIdBeBindModule: []int{1, 2, 3},
	//})
	//fmt.Println("result error :", err)
	//
	//idTest := bson.NewObjectId()
	//idTest1 := bson.NewObjectId()
	//err = r.C.Insert(map[string]interface{}{
	//	"test": []bson.ObjectId{idTest1},
	//	"_id":  idTest,
	//})
	//fmt.Println("result id :", idTest)
	//
	//ids := bson.NewObjectId()
	//err = r.C.Insert(BindFileIDMgo{
	//	Test: []bson.ObjectId{ids},
	//})

	//twoLevelArrMacthFind(r)

	usePipeRequiredArray(r)
	//testObjectJson()
}

// 匹配两层数组的方式
// 删除数组元素
func twoLevelArrMacthFind(r Repository) {
	idTest123 := bson.NewObjectId()
	err := r.C.Insert(BindFileIDMgo{
		Id:                 idTest123,
		ProductId:          123,
		FileIdBeBindModule: []int{1, 2, 3},
		ArrTest: []TwoParam{
			{Test1: 1, Test2: bson.NewObjectId()},
			{Test1: 12, Test2: bson.NewObjectId()},
		},
	})

	var testData interface{}
	err = r.C.Find(map[string]interface{}{
		"_id": idTest123,
		"testArr": map[string]interface{}{
			"$elemMatch": map[string]interface{}{"test1": 1},
		},
	}).One(&testData)

	fmt.Println("error :", err)
	fmt.Println("result 123 :", testData)

	// 删除数组元素
	r.C.Update(map[string]interface{}{
		"_id": idTest123},
		map[string]interface{}{
			"$pull": map[string]interface{}{
				"testArr": map[string]interface{}{
					"test1": 1}}})
	fmt.Println("delete param result :", idTest123, err)

	err = r.C.Update(map[string]interface{}{
		"_id": idTest123,
	}, map[string]interface{}{
		"$push": map[string]interface{}{
			"testArr": TwoParam{Test1: 1123,
				Test2: bson.NewObjectId()}}})
}

// 匹配单层数组的方式
func oneLevelArrMatchFind(r Repository) {
	var bind BindFileIDMgo
	err := r.C.Find(map[string]interface{}{
		"test": map[string]interface{}{
			"$elemMatch": map[string]interface{}{"$eq": "123"},
		},
	}).One(&bind)
	fmt.Println("error test :", bind, err)
}

// 删除数组中的元素
func deleteArrayParam(r Repository) {
	idTest := bson.NewObjectId()
	err := r.C.Update(map[string]interface{}{
		"_id": idTest,
	}, map[string]interface{}{
		"$pull": map[string]interface{}{
			"test": "123",
		},
	})
	fmt.Println("error :", err)
}

// 确保插入不重复的参数
func noRepeatInsertParam(r Repository) {
	err := r.C.Update(map[string]interface{}{
		"_id": bson.ObjectIdHex("5bebe0cf7f45aa3270c9e532"),
	}, map[string]interface{}{
		"$addToSet": map[string]string{
			"test": "1234567890",
		},
	})
	fmt.Println("error :", err)
}

// 聚合的复杂方法
func usePipeRequiredArray(r Repository) {
	idTest123 := bson.NewObjectId()
	err := r.C.Insert(BindFileIDMgo{
		Id:                 idTest123,
		ProductId:          123,
		FileIdBeBindModule: []int{1, 2, 3},
		ArrTest: []TwoParam{
			{Test1: 1, Test2: bson.NewObjectId()},
			{Test1: 12, Test2: bson.NewObjectId()},
		},
	})
	fmt.Println("error1 :", err)

	pipe := r.C.Pipe([]bson.M{
		{"$match": bson.M{"_id": idTest123, "product_id": 123}},
		{"$unwind": "$testArr"},
		{"$skip": 1},
		{"$limit": 1},
		{"$project": bson.M{"product_id": 1, "one": "$testArr.test1"}},
	})
	resp := bson.M{}
	err = pipe.One(&resp)
	fmt.Println("error is :", err, resp)
}

// 使用MongoDB中的聚合操作，计算MongoDB嵌套数组的大小
func requiredArraySize(r Repository) {
	idTest123 := bson.NewObjectId()
	pipe := r.C.Pipe([]bson.M{
		{"$match": bson.M{"_id": idTest123}},
		{"$project": bson.M{"count": bson.M{"$size": "$fileid_bind_module"}}},
	})
	resp := bson.M{}
	err := pipe.One(&resp)
	fmt.Println("error is :", err, resp)
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
