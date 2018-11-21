package main

import (
	"fmt"
)

func Test() {
	ts := make([]int, 5, 5)
	ts1 := ts

	ts1 = ts1[2:]

	fmt.Println(ts)
	fmt.Println(ts1)
	fmt.Println(ts1[1])

}

// 反转数组
func reverseArr(arr []int) {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func testSlice(arr []int) {
	arrTmp := arr
	arr = append(arr[:1], arr[2:]...)

	fmt.Println(arrTmp)
	fmt.Println(arr)
	fmt.Println(arr[1:])
}

// 数组一边减少一边range
func testRangeSlice(arr []int) {
	fmt.Printf("%p \n", &arr)
	fmt.Println(cap(arr), len(arr), arr)

	for index, a := range arr {
		if index == 2 {
			arr = append(arr[:2], arr[3:]...)
			fmt.Println(cap(arr), len(arr), arr)
		}
		fmt.Println("index", index, a)
	}
	fmt.Println(cap(arr), len(arr), arr)
	fmt.Printf("%p \n", &arr)
}

func testAppend() {
	s := make([]int, 0, 1)
	oldcap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newcap := cap(s); oldcap < newcap {
			fmt.Printf("oldcap %d ===> newcap %d\n", oldcap, newcap)
			oldcap = newcap
		}
	}
}

func main() {
	//v := []int{1, 2, 3}
	//for i := range v {
	//	v = append(v, i)
	//}
	//
	//fmt.Println(v)

	//testAppend()

	testRangeSlice([]int{1, 2, 3, 4, 5, 6})

}

//oldcap 1 ===> newcap 2
//oldcap 2 ===> newcap 4
//oldcap 4 ===> newcap 8
//oldcap 8 ===> newcap 16
//oldcap 16 ===> newcap 32

//func main() {
//	//Test()
//
//	testArr := []int{1, 2, 3, 4, 5}
//	//reverseArr(testArr)
//	//fmt.Println(testArr)
//	//testSlice(testArr)
//
//	testRangeSlice(testArr)
//}
