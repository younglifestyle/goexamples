package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	head, tail := 0, len(arr)-1
	qviot := arr[head]

	for i := 1; i < tail; {
		if arr[i] > qviot {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}

	quickSort(arr[:head])
	quickSort(arr[head+1:])
}

func main() {
	a1 := []int{1, 0, 45, 34, 65, 0, 12, 43, 52}
	quickSort(a1)

	fmt.Println("快速排序结果：")
	fmt.Println(a1)
}
