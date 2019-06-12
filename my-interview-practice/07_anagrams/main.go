package main

import (
	"fmt"
	"sort"
	"strings"
)

// anagram 是指，两个字符串中，如果字符的数量相同，则是 anagram 如果不同，则不是
// 对应leetCode 242

// 使用map对str1和str2中的不同字节进行计数
// 对 str1 进行加
// 对 str2 进行减
func anagram(s, t string) (res bool) {

	m := make(map[rune]int)
	runes1 := []rune(s)
	runes2 := []rune(t)

	length := len(runes1)
	if length != len(runes2) {
		return false
	}

	for i := 0; i < length; i++ {
		m[runes1[i]] += 1
		m[runes2[i]] -= 1
	}

	for _, c := range m {
		if c != 0 {
			return false
		}
	}

	return true
}

// 使用排序确定字符串是否一致
func sortAnagram(s, t string) (result bool) {

	length := len(s)
	if length != len(t) {
		return false
	}

	sortFunc := func(s string) string {
		strSplit := strings.Split(s, "")
		sort.Strings(strSplit)
		return strings.Join(strSplit, "")
	}

	return sortFunc(s) == sortFunc(t)
}

func main() {
	fmt.Println(anagram("anagram", "nagaram"))

	fmt.Println(sortAnagram("anagram", "nagaram"))
}
