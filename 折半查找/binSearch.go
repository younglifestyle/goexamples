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
			return mid
		} else if needFindData > list[mid] {
			lowIndex = mid + 1
		} else {
			highIndex = mid - 1
		}
	}

	return mid
}

func directSelectAlgo(arr []int) {

	for i := 0; i < len(arr)-1; i++ {

		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}

	fmt.Println(arr)
}

func main() {
	a := []int{1, 4, 5, 6, 8, 12, 43, 52}

	fmt.Println("index :", binSearch(a, 52))

	a1 := []int{1, 45, 34, 65, 0, 12, 43, 52}
	directSelectAlgo(a1)

}
