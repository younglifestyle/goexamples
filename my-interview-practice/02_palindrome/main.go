package main

import (
	"fmt"
)

// 检查数字是否是回文数
func checkNum2(num int) bool {

	if num < 0 || (num != 0 && num%10 == 0) {
		return false
	}

	return num == reverseInt(num)
}

func reverseInt(num int) int {
	res := 0

	for ; num != 0; num = num / 10 {
		res = res*10 + num%10
	}

	return res
}

// 检查数字是否是回文数  类似于先转换成字符串再做比较
func checkNum(num int) bool {

	if num < 0 || (num != 0 && num%10 == 0) {
		return false
	}

	numByte := []byte{}

	for {
		theNum := num % 10
		num = num / 10
		if theNum == 0 && num == 0 {
			break
		}
		numByte = append(numByte, byte(theNum))
	}

	for i, j := 0, len(numByte)-1; i < j; i, j = i+1, j-1 {
		if numByte[i] != numByte[j] {
			return false
		}
	}

	return true
}

// 检查字符串是否是回文数
func checkPalindRome(str string) bool {
	if str == "" {
		return false
	}

	toBytes := []rune(str)
	for i, j := 0, len(toBytes)-1; i < j; i, j = i+1, j-1 {
		if toBytes[i] != toBytes[j] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkPalindRome("中国中"))
	fmt.Println(checkPalindRome("123321"))
	fmt.Println(checkPalindRome(""))

	fmt.Println(checkNum(1101))
	fmt.Println(checkNum2(-1911))
}
