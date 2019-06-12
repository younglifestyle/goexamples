package main

import "fmt"

// 输入一个整数n
// 输出用 # 和 <space> 表示的阶梯
// 比如 输入3
// '#  '
// "## "
// "###"

func printfStep(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j <= i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
func prinrf1(n int) {
	prinrfStep1(n, 0)
}

// row, col 行，列
func prinrfStep1(row, col int) {

	if row == 0 {
		return
	}

	prinrfStep1(row-1, col+1)
	for index := 0; index < row; index++ {
		fmt.Print("#")
	}
	for index := 0; index < col; index++ {
		fmt.Print(" ")
	}
	fmt.Println()
}

func main() {
	printfStep(3)
	fmt.Printf("\n")
	printfStep(5)

	prinrf1(3)
	prinrf1(2)
}
