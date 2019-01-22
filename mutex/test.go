package main

import (
	"fmt"
	"sync"
	"time"
)

var l sync.Mutex

func main() {
	l.Lock()
	go f()
	time.Sleep(3000 * time.Millisecond)
	l.Unlock()

	fmt.Println("finish")
	for {
		time.Sleep(3000 * time.Millisecond)
	}

}

func f() {
	l.Lock()
	fmt.Println("i am going")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("hello world")
	l.Unlock()
}
