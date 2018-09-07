package main

import "fmt"

func directselectAlgo(a []int) {

	for index, _ := range a {
		min := index

		for in, _ := range a {

			if a[min] < a[in] {
				min = in
			}

			fmt.Println("li ", min)
		}
		fmt.Println("min ", index, a[index], min, a[min])
		if min != index {
			t := a[min]
			a[min] = a[index]
			a[index] = t
		}
	}

}

func main() {
	in := []int{1, 5, 142, 34, 231, 2412}

	directselectAlgo(in)

	fmt.Println("arrary :", in)
}
