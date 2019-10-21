package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string
	fmt.Println("string :", str, unsafe.Sizeof(str))

	var intr interface{}
	fmt.Println("interface :", intr)

	var slic []int
	fmt.Println("slice :", slic, unsafe.Sizeof(slic))

}
