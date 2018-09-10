package main

import "fmt"

func Test() {
	ts := make([]int, 5, 5)
	ts1 := ts

	ts1 = ts1[2:]

	fmt.Println(ts)
	fmt.Println(ts1)
	fmt.Println(ts1[1])
}

func main() {
	Test()
}
