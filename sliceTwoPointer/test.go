package main

import "fmt"

func FuncSlice(s []int, t int) {
	s[0]++
	s = append(s, t)
	s[0]++
}

func testSlice() {
	ts := "12321421421521"

	fmt.Println(ts[:len(ts)-2])
	fmt.Printf("%s \n", string(ts[len(ts)-2:]))
}

func main() {
	a := []int{0, 1, 2, 3}
	FuncSlice(a, 4)
	fmt.Println(a)

	testSlice()

}
