package main

import "fmt"

func main() {
	var ch chan int
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

	select {}
}
