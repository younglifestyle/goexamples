package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// 短链字符: a-z 0-9 A-Z
var Codes = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5',
	'6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z'}

func GenerateShortLink(longURL string) string {
	// 生成4个短链串
	loopNum := 4
	var urls []string
	var i int

	// long url转md5
	md5URL := md5.New()
	// 此处的"salt": 自定义字符串,防止算法泄漏
	_, err := io.WriteString(md5URL, "salt"+longURL)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	md5Byte := md5URL.Sum(nil)
	md5Str := fmt.Sprintf("%x", md5Byte)

	for i = 0; i < loopNum; i++ {
		// 每8个字符是一个组
		each8BitsStr := md5Str[i*8 : i*8+8]
		// 将一组串转成16进制数字
		val, err := strconv.ParseInt(fmt.Sprintf("%s", each8BitsStr), 16, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 获得一个新的短链
		urls = append(urls, genShortURL(val, i))
	}

	// 下面从上面的4个短链中随机一个作为当前的短链
	if len(urls) == 0 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := r.Intn(len(urls))
	return urls[idx]
}

// 32位的md5串中, 每8个字符作为一组, 计算得到一个6个字符的短链
func genShortURL(every8Chars int64, idx int) string {
	// 首先every8Chars是md5中的一组子串, 8个字符
	// 所有的数据支取后面30bits
	base := every8Chars & 0x3FFFFFFF
	// 数组保存6个字符
	result := make([]byte, 6)
	// 下面生成6个字符
	for j := 0; j < 6; j++ {
		// 0x0000003D = 61, 所以0x0000003D & out保证生成的数载0~61, 所以其实就是Codes的所有下标
		idx := 0x0000003D & base
		// 获取这个idx下标的字符
		result[j] = Codes[int(idx)]
		// 继续处理后面的bits
		base = base >> 5
	}
	return fmt.Sprintf("%s", result)
}

func main() {
	fmt.Println(GenerateShortLink("http://www.baidu.com"))
}
