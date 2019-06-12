package main

import (
	"fmt"
	"goexamples/my-interview-practice"
	"strings"
)

// 将一个句子中的每一个词的首字母大写

// 我的这个处理方法，是建立在句子是正常空格，hello, world. 要是碰上hello,world.就不行了
func sentenceCapitalization(s string) (str string) {

	strSplit := strings.Split(s, " ")

	for index, _ := range strSplit {
		if strSplit[index][0] >= 'a' && strSplit[index][0] <= 'z' {
			toBytes := []byte(strSplit[index])
			toBytes[0] = toBytes[0] - 32
			strSplit[index] = utils.ToString(toBytes)
		}
	}

	return strings.Join(strSplit, " ")
}

func sc(str string) string {
	sSplit := strings.Split(str, "")
	for index, _ := range sSplit {
		if index == 0 || sSplit[index-1] == " " || sSplit[index-1] == "," {
			if sSplit[index] >= "a" && sSplit[index] <= "z" {
				sSplit[index] = strings.ToUpper(sSplit[index])
			}
		}
	}

	return strings.Join(sSplit, "")
}

func main() {
	fmt.Println(sentenceCapitalization("I'am super man, yeah"))
	fmt.Println(sc("I'am super man,yeah"))
}
