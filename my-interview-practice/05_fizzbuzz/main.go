package main

import "fmt"

// fizzbuzz问题：
// 输入一个整数
// 找出其中的fuzz数，3的倍数
// 找出其中的buzz数,5的倍数

// 号称是面试中，经常出现的最简单题
// 对应leetcode 412

func fuzzBuzz(num int) (fuzzByte []byte, buzzByte []byte) {

	fuzzByte = []byte{}
	buzzByte = []byte{}

	for ; num != 0; num = num / 10 {
		tmp := num % 10
		if tmp%3 == 0 {
			fuzzByte = append(fuzzByte, byte(tmp))
		} else if tmp%5 == 0 {
			buzzByte = append(buzzByte, byte(tmp))
		}
	}

	return fuzzByte, buzzByte
}

func fuzzbuzz(num int) {
	for i := 1; i <= num; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {

	fuzzByte, buzzByte := fuzzBuzz(654309876543123)
	fmt.Println(fuzzByte, buzzByte)

	fuzzbuzz(18)
}
