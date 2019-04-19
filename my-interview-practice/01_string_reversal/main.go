package main

import (
	"fmt"
	"goexamples/my-interview-practice"
)

// 遍历字符串
// golang 中没有build-in的例如js中的 reverse()的方法对数组进行反转，所以得手写一个反转数组的方法
// 为什么这里使用了 rune ,而不是 byte ?
// byte 是 8bit
// rune 是 32bit
// 在utf-8编码中，对于 中文字符 而言，一个字符占 3个字节, 使用 byte 是放不下的
// 常见的 range 也是对str进行了隐式的 unicode 解码, 而 str[i] 并不一定和我们看到的字符串对应
// 同理，如果只是序列化和反序列化，可以通过byte进行操作，但是如果涉及字符串中的反转，截断等操作，则使用rune
func reverseByRune(reversalStr string) string {

	toBytes := []rune(reversalStr)

	for i, j := 0, len(toBytes)-1; i < j; i, j = i+1, j-1 {

		toBytes[i], toBytes[j] = toBytes[j], toBytes[i]
	}

	return string(toBytes)
}

func reverseByByte(reversalStr string) string {

	// 不能使用这种方式，下面替换语句会报错
	//toBytes := utils.ToBytes(reversalStr)
	toBytes := []byte(reversalStr)

	for i, j := 0, len(toBytes)-1; i < j; i, j = i+1, j-1 {

		toBytes[i], toBytes[j] = toBytes[j], toBytes[i]
	}

	return utils.ToString(toBytes)
}

func main() {

	fmt.Println(reverseByByte("test mian"))
	fmt.Println(reverseByRune("hello 中国"))
}
