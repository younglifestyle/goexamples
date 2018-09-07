package main

import (
	"fmt"
	"sync"
)

type Test struct {
	Name string
	Age  int
}

func TestMap() {

	mapTest := make(map[int]interface{})

	mapTest[1] = Test{
		Name: "test",
		Age:  18,
	}

	test, ok := mapTest[1]
	if ok {
		k, v := test.(Test)
		if v {
			fmt.Println("断言成功 :", k)
		}
	}
}

type userInfo struct {
	Name string
	Age  int
}

var m sync.Map

func TestSyncMap() {

	vv, ok := m.LoadOrStore("1", "one")
	fmt.Println(vv, ok) //one false

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	vv, ok = m.LoadOrStore("1", "oneone")
	fmt.Println(vv, ok) //one true

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	m.Store("1", "oneone")
	vv, ok = m.Load("1")
	fmt.Println(vv, ok) // oneone true

	m.Store("2", "two")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v, "range")
		return true
	})

	m.Delete("1")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v, "range Delete")
		return true
	})

	map1 := make(map[string]userInfo)
	var user1 userInfo
	user1.Name = "ChamPly"
	user1.Age = 24
	map1["user1"] = user1

	var user2 userInfo
	user2.Name = "Tom"
	user2.Age = 18
	m.Store("map_test", map1)

	mapValue, _ := m.Load("map_test")

	for k, v := range mapValue.(map[string]userInfo) {
		fmt.Println(k, v)
		fmt.Println("name:", v.Name)
	}
}

func main() {

	//TestMap()

	TestSyncMap()

}
