package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

var ALPHABET = strings.Split("adefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")

func main() {
	fmt.Println(GetShortUrl(10000))

	randNo := int64(random(1, 99))*int64(math.Pow(10, 9)) + 1000

	s := strconv.Itoa(int(randNo))
	fmt.Println(s)
	fmt.Println(GetShortUrl(randNo))
}

func GetShortUrl(id int64) string {
	indexAry := Encode62(id)
	return GetString62(indexAry)
}

// 转换成62进制
func Encode62(id int64) []int64 {
	indexAry := []int64{}
	base := int64(len(ALPHABET))

	for id > 0 { // i < 0 时,说明已经除尽了,已经到最高位,数字位已经没有了
		remainder := id % base
		indexAry = append(indexAry, remainder)
		id = id / base
	}

	return indexAry
}

//  输出字符串, 长度不一定为6
func GetString62(indexAry []int64) string {
	result := ""

	for val := range indexAry {
		result = result + ALPHABET[val]
	}

	return reverseString(result)
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
