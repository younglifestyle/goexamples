package main

import "fmt"

// 区分大小写，私有，公有变量
type test struct {
	Value int
	Data  int
	eTime int
}

func (ts test) Print() {
	fmt.Print(ts.Value, " ")
}

func main() {
	var ts test

	ts.Value = 1
	ts.Print()
}

// 包管理演示一下