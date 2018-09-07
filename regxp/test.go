package main

import (
	"fmt"
	"regexp"
	"unicode"
)

//var nicknamePattern = `[\p{Han}_a-zA-Z0-9_]{4,10}`
var nameRegexp = regexp.MustCompile(`^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*$`)

var ipRegexp = regexp.MustCompile(`((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`)

var yingwen = regexp.MustCompile(`^[a-zA-Z_0-9]+$`)
var version = regexp.MustCompile(`^([1-9]\d|[1-9])(\.([1-9]\d|\d)){1,2}$`)

var reg = regexp.MustCompile(`(((\d{4})(0[13578]|10|12)(0[1-9]|[12][0-9]|3[01]))|((\d{4})(0[469]|11)(0[1-9]|[12][0-9]|30))|((\d{4})(02)(0[1-9]|1[0-9]|2[0-8]))|([0-9][0-9][02468]40229)|([0-9][0-9][02468]80229)|([0-9][0-9][13579]20229)|([0-9][0-9][13579]60229)|([0-9][0-9][02468]00229))`)
var date = regexp.MustCompile(`(20\d{2})(\d{2})(\d{2})`)

// (\d{1,3}\.){3}\d{1,3}
// `^[a-zA-Z][a-zA-Z0-9_]{4,15}$`  `^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*$`
//`^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+){1,20}$`

//[a-zA-Z\d\u4e00-\u9fa5]  [A-Za-z0-9_\-\u4e00-\u9fa5]+
func main() {
	str := "1231"
	//var hzRegexp = regexp.MustCompile(nameRegexp)
	fmt.Println(nameRegexp.MatchString(str))
	fmt.Println(IsChineseChar(str))

	strs := "192.168.123.1232"
	fmt.Println("sad", ipRegexp.MatchString(strs))

	str1 := "12321312"
	fmt.Println("yes", yingwen.MatchString(str1))

	str2 := "12.2"
	fmt.Println("yes2", version.MatchString(str2))

	strTest := "20120229-j"
	fmt.Println("日期格式：", reg.MatchString(strTest))
	//fmt.Println("日期格式1：", reg.FindAllSubmatch([]byte(strTest), -1))

	submatch := reg.FindAllSubmatch([]byte(strTest), -1)
	for _, val := range submatch {
		fmt.Println(string(val[0]))
	}

	strTest1 := "20120230"
	fmt.Println("日期格式：", date.MatchString(strTest1))

	//bytes, e := json.Marshal(nil)
	//fmt.Println(string(bytes), e)
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
