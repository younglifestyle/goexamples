package main

import (
	"fmt"
	"time"
)

func main() {
	delayChan := NewDelayChan(time.Second * 2)

	<-delayChan
	fmt.Println("delay...")
}

// NewDelayChan 发送延时信号
func NewDelayChan(t time.Duration) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		time.Sleep(t)
		close(c)
	}()
	return c
}
