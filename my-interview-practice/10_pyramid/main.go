package _10_pyramid_

import "fmt"

// 和上一个阶梯类似，输出金字塔型的#
// 例如输出3
// "  #  "
// " ### "
// "#####"

func pyramid(n int) {
	printPy(n*2-1, 0)
}

func printPy(i, j int) {
	if i <= 0 {
		return
	}

	// 输出少两个# ，多个<space>
	printPy(i-2, j+1)
	for index := 0; index < j; index++ {
		fmt.Print(" ")
	}
	for index := 0; index < i; index++ {
		fmt.Print("#")
	}
	for index := 0; index < j; index++ {
		fmt.Print(" ")
	}
	fmt.Println()
}
