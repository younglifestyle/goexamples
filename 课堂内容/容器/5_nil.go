package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var str string
	fmt.Println("string :", str, unsafe.Sizeof(str))
	fmt.Println("len string :", str, len(str))

	var intr interface{}
	fmt.Println("interface :", intr)

	var slic []int
	fmt.Println("slice :", slic, unsafe.Sizeof(slic))

	fmt.Println(math.Pow(10, 9))
}
