package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go thread()
	fmt.Printf("Main pid: %d\n", os.Getpid())
	time.Sleep(100 * time.Second)
}

func thread() {
	fmt.Printf("Child pid: %d", os.Getpid())
	time.Sleep(100 * time.Second)
}
