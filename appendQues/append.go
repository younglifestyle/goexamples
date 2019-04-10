package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[:5]
	fmt.Println(s2)
}
