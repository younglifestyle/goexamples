package main

import (
	"fmt"
)

func binSearch(list []int, needFindData int) (index int) {

	lowIndex := 0
	highIndex := len(list) - 1

	var mid int

	for lowIndex <= highIndex {
		mid = (lowIndex + highIndex) / 2 // 此处有可能溢出

		if list[mid] == needFindData {
			return mid + 1
		} else if needFindData > list[mid] {
			lowIndex = mid + 1
		} else {
			highIndex = mid - 1
		}
	}

	return 0
}

func main() {
	a := []int{1, 4, 5, 6, 8, 12, 43, 52}
	fmt.Println("index :", binSearch(a, 5))
}
