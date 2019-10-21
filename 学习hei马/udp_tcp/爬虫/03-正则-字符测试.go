package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "abc a7c mfc cat aMc azc cba"
	// 解析、编译正则表达式
	//ret := regexp.MustCompile(`a.c`)  	// `` : 表示使用原生字符串
	ret := regexp.MustCompile(`a[^0-9a-z]c`) // `` : 表示使用原生字符串

	// 提取需要信息
	alls := ret.FindAllStringSubmatch(str, -1)
	fmt.Println("alls:", alls)

}
