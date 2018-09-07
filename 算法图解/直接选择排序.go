package main

import "fmt"

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
}

func main() {
	a1 := []int{1, 45, 34, 65, 0, 12, 43, 52}
	directSelectAlgo(a1)

	fmt.Println("直接选择排序结果：")
	fmt.Println(a1)
}
