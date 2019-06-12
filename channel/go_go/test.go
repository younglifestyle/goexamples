package main

import (
	"fmt"
	"time"
)

var tCh chan struct{}

func main() {
	tCh = make(chan struct{})

	go thread()

	<-tCh
}

func thread() {

	time.Sleep(time.Second)

	fmt.Println("go...")

	tCh <- struct{}{}
}
