package main

import "fmt"

func channelTest() {
	c := make(chan int)
	c <- 1
	c <- 2
	n := <-c

	fmt.Println(n)
}

// 改进

func main() {
	channelTest()
}
