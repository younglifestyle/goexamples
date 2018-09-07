package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdef"))
}

func lengthOfNonRepeatingSubStr(s string) int {

	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	var subString string

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
			subString = s[start : start+maxLength]
		}

		lastOccurred[ch] = i
	}

	fmt.Println(subString)

	return maxLength
}

func findNoCopySubString(str string) string {

	if str == "" {
		return ""
	}

	subStrArr := []string{}
	subStrArrBack := []string{""}
	strMap := make(map[string]string)
	strMapBack := make(map[string]string)

	for {
		oneStr := string(str[0])
		str = str[1:]

		_, found := strMap[oneStr]
		if !found {
			strMap[oneStr] = oneStr
			subStrArr = append(subStrArr, oneStr)
		} else {
			if len(str) > len(strMap) &&
				len(subStrArrBack) < len(subStrArr) {
				strMap = make(map[string]string)
				subStrArr = []string{}

				strMap[oneStr] = oneStr
				subStrArr = append(subStrArr, oneStr)
			} else {
				strMapBack = strMap
				subStrArrBack = subStrArr
			}
		}

		if len(str) == 0 {
			break
		}
	}

	fmt.Println(strMapBack, subStrArrBack)

	subString := strings.Replace(strings.Trim(fmt.Sprint(subStrArr), "[]"),
		" ", "", -1)

	return subString
}
