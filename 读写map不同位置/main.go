package main

import "fmt"

func main() {
	var deviceOnline bool

	testMap := make(map[int]int)

	testMap[1] = 1
	testMap[2] = 1

	// 写不同位置的map也会协程冲突
	//go func() {
	//	i := 0
	//	for true {
	//		testMap[1] = i
	//		i++
	//	}
	//}()
	//
	//go func() {
	//	i := 0
	//	for true {
	//		testMap[2] = i
	//		i++
	//	}
	//}()

	// 读不同位置，暂时未报错
	//go func() {
	//	for true {
	//		fmt.Println(testMap[1])
	//	}
	//}()
	//
	//go func() {
	//	for true {
	//		fmt.Println(testMap[2])
	//	}
	//}()

	go func() {
		i := 0
		for {
			deviceOnline = true
			i++
			if i %2 == 0 {
				deviceOnline = false
			}
		}
	}()

	go func() {
		for {
			fmt.Println(deviceOnline)
		}
	}()

	select {}
}
