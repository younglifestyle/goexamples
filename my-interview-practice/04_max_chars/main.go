package main

import (
	"fmt"
	"goexamples/my-interview-practice"
)

//  max chars : 找到一个字符串中，数量最多的那一个字符
//  遇到字符串相关的算法，首先要问自己3个问题：
// 1. 字符串中最常见的字符是什么？
// 2. 输入和输出的字符串有相同的字符吗？
// 3. 给出的字符串中有重复的字符吗？

// 可以用此方法去除字符串中重复的char

func maxChars(str string) (indexChar string, cnt int) {

	calcCharMap := make(map[byte]int)
	toBytes := utils.ToBytes(str)
	if len(toBytes) == 0 {
		return "", 0
	}

	for _, char := range toBytes {
		calcCharMap[char] = calcCharMap[char] + 1
		if calcCharMap[char] > cnt {
			cnt = calcCharMap[char]
			indexChar = string(char)
		}
	}

	return
}

func main() {

	indexChar, cnt := maxChars("1222223aa vfdvdddddd")
	fmt.Println(indexChar, cnt)
}
