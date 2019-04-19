package main

import "fmt"

// 还有一种是使用回文法，将整数先转化为[]byte，然后倒置

func reverseInt(num int) (newNum int) {

	for ; num != 0; num = num / 10 {
		newNum = newNum*10 + num%10
	}

	return
}

func P2(num int) (newNum int) {
	for num != 0 {
		temp := newNum*10 + num%10
		if temp/10 != newNum {
			return 0
		}
		newNum, num = temp, num/10
	}

	return
}

func main() {
	fmt.Println(P2(-123142))
	fmt.Println(reverseInt(-123142))
}
