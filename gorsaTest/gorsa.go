package main

import (
	"errors"
	"log"

	"encoding/base64"

	"fmt"

	"bytes"

	"github.com/wenzhenxi/gorsa"
)

var Pubkey = `-----BEGIN -----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDLftGiduO2ZS1m4dvQXYUuFak9
Ryo7lRrftsKxpv58+kY1UqgymbwPUW12/dXhSYEE7NNUzE9uB39qT4twALb7yIFO
QFmlsU0ymoiNCkDzUlECPABmeo5MjM5T+L4FEh53oRbgR/AotEQJw3/uIVGs0SFd
XI1rb4kX2r/ZmpbcVQIDAQAB
-----END -----
`

var Pirvatekey = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDLftGiduO2ZS1m4dvQXYUuFak9Ryo7lRrftsKxpv58+kY1Uqgy
mbwPUW12/dXhSYEE7NNUzE9uB39qT4twALb7yIFOQFmlsU0ymoiNCkDzUlECPABm
eo5MjM5T+L4FEh53oRbgR/AotEQJw3/uIVGs0SFdXI1rb4kX2r/ZmpbcVQIDAQAB
AoGADLir4E0wZRmkIfdip38BMWVXRCZrxHfIy02AlFyNMkDvHKxDKY6kzAxaHIM3
2LKgpy8q8vUmzgkX9Qxt+h2BUtWK99PfmMrYT2+bEP9vBTFogXln10BYsyQcpYO8
Gnsk27hUaipCvRXX6QFGoLsV5yKD/nlTbup2/2ReCKqdSo0CQQDQ81xDJzYtDYEC
tUOmb5mmWn1qCa2FYfgoIDoQQDU8czWWoAg6ZER+5UkFTAfbx/vxj50l9YxYUDTU
VSLXWznDAkEA+VD/3t2KK4OY2rcRf09XxTnCf4bVcGVpBmdkTb2YZSzcHdN6WJi9
3nRB0yh24RXUlqQ/OYksJAl4Rz/Y5roYBwJAQm/+uANxwGV4zcmg1vzJGBHRqrOl
DrIV23xTufMQekYPlfMQarCS7t4sl5iTLxipTSdiyj0HANWP1quzRlJlTQJAKG1l
ADuGQyYUrCqRUMaJ4fZKvqkbhR08mYg8cIq04nsSuldneGRULXVGkzn1hOwoS8EY
a3j9yl4qvcrxngBBEwJAXbAPYmPlbQ0GWB0s0hYJjs0sQqeeB3qZGGIIi7sC4G+w
tWMrWFEMrB/CjFZsG3UiZ8OMA01hpQlJ33ZBagy7DQ==
-----END RSA PRIVATE KEY-----
`

func main() {
	// 公钥加密私钥解密
	if err := applyPubEPriD(); err != nil {
		log.Println(err)
	}
	// 公钥解密私钥加密
	if err := applyPriEPubD(); err != nil {
		log.Println(err)
	}

	b := string(`hello,worldhello,worldhello,worldhello,worldhello,worldhello,world
hello,worldhello,worldhello,worldhello,worldhello,worldhello,worldhello,worldhello,worldhello,world
hello,worldhello,worldhello,worldhello,worldhello,`)

	encrypt117 := divideEncrypt117(b)
	divideEncrypt117(string(encrypt117))

	//if err := doublePrivEntry(); err != nil {
	//	log.Println("error:", err)
	//}

	//pubdecrypt, err := gorsa.PublicDecrypt(`gEmpUAv7QH4XioRwDUreQsxHUrM0Ucja5N6hPyI6UtbPIfzFrZfD7BR+DAGio3QQ32mll2tZwVqATaRtACopujGFbY0fO802EPa/g/AAGlhMYqEqDXbnGWRPg0uY1HRy6KcgHVX0V6KaTRwzvff/6Es+d5KPHsUP91NAEgYcyAwiOxTtsZdVXQCoZZmiAB2V93u/QIE5p6UNPqpdvvJH9iQ+CcTQCE+HNyaEN1xrbaGeyUKZU4cTtsJ5AEXZazcLZto6eKv8Zcg/r1JC3TbsYwGqayqRbxYp6jWMgYn76fVNl29gGY800ZLsbwa8gCBmU42zRDR6ZfhNYG8UaFr3PR1s6SC+IDjgx3HpQyivOd16QRwdDA1o6Kw1HhrsJgI4cAzsXj0yNafx3+7gke6hqouIUfaoq1mvGGqrSTZaJ/TJJ9KxWp0Z9MsT3O51Xp1RZNb9MAgeDRSWZUBo80s/jTqXW5DfgpW38sSOy1ffSGsXQ7XWoYbGeNVklDIV8SzPDs6RmaOEuAqoy6CHOMPorTxr9qwYTYtZg4eBTYWb99MN+Tfhss6Ne/K5jJj9XomayAi4xa7Dk+uIEqKrQnDhWp5pTJFOAasGoB7zXED90ZgDxk9TBnCIlU6oO6kmk3eWpEH417WE6++FlM1yxEGPXtvFClmSu4wZIktrtRptW8KbO1+mZHsGLi4KBmNZqPB7FP6VzJ8DI0Ef8EmOpzlx0ub5RXFfiml8uTAWJbS7d2HWPdYyMqrRF29in+ujTYtkZtp5L3Dpttf3xH6VTXmZNsEarfwnlObX+z2Z+RxIVi3LXszrTQzbqFJGTUXP4D/hDPxmkYAUINDPz+JwCp158D2/MWS2kFqZ7dMGSCEa0QUvE/WTkxpk6cdVoGRwaQxeZsq0fglh+0LewHaMXjkC0ByAu9iZEHVwm8BXLoFp0oBIknvWW4zoO22/BYsYP+wSN7fKow5+S+cCb4NE3n021ZUayKi7Fsey5Qagknen3vyseuxkX8RlSYu8ce/rdz2NiYFir0efbOWIQBOOGGK8YjMehUt7PiPFkLOHFtZwejwd4SWfwocPi0Oje0YOeRAPuUidwRvNlFTgi7oyq4TDEIogij3j2bYiCXO4WCku3KpEIVVa9szE02GqFKMRDoGXMPMtqkJGIHhSbHxxT6eT/CgqvbGtx4F9V61pSn2WPyAj5JzEvnhOQDXcgBEaiSJbP/A9mCoRQnO6kcZd7xMETvMAWQeMqJjg4m80vv155MZQxVZp8T3XOUpmc+RumlkOut87PF1Tx6gAZu/stW3FCmAwWKxBmH7Qz1rhRrhWkB3IzIJvYv7pNJfWmR169I6j3hSEg8ihUzQw/d4iQ+8Xlz/6NSbnje7AIf6fNPcWfD3zwZ7+KULcxQQROrrqN/StqnNMBUYtFR1jrGiG/sruxQgxGCuyFTynZ293Ut/SGWvZ/GYhRp5ziy+Rog19Pubt26t3XFpoHyFxXj0DnVIfi13Dd5R3V7P0FoJH2QakfAvrS9/VLELkWO4G0M1RT9oO`,
	//	Pubkey)
	//if err != nil {
	//	log.Println("error1:", err)
	//	return
	//}

	//toString := base64.StdEncoding.EncodeToString([]byte(pubdecrypt))
	//pubdecrypt1, err := gorsa.PublicDecrypt(toString,
	//	Pubkey)
	//if err != nil {
	//	log.Println("error2:", err)
	//	return
	//}
	//
	//fmt.Println("bytes :", pubdecrypt1)
}

//BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func divideEncrypt117(str string) []byte {

	var priEncryptData string
	var priEnc []byte

	lens := len(str)
	oneBlockSize := lens / 117

	for i := 0; i < oneBlockSize; i++ {

		prienctypt, err := gorsa.PriKeyEncrypt(str[i*117:(117+i*117)], Pirvatekey)
		if err != nil {
			fmt.Println("error 1", err)
			return nil
		}
		byteStr, _ := base64.StdEncoding.DecodeString(prienctypt)
		priEncryptData += string(byteStr)

		priEnc = BytesCombine(priEnc, byteStr)
	}

	fmt.Println("byte :", priEnc)

	prienctypt, err := gorsa.PriKeyEncrypt(str[(oneBlockSize*117):], Pirvatekey)
	if err != nil {
		fmt.Println("error 2", err)
		return nil
	}

	byteStr, _ := base64.StdEncoding.DecodeString(prienctypt)
	priEncryptData += string(byteStr)

	priEnc = BytesCombine(priEnc, byteStr)

	fmt.Println(base64.StdEncoding.EncodeToString(priEnc))

	return priEnc
}

// 库中RSA加密使用分段加密，即RSA
func doublePrivEntry() error {

	prienctypt, err := gorsa.PriKeyEncrypt(`hello world`, Pirvatekey)
	if err != nil {
		return err
	}

	fmt.Println("密文1：", prienctypt)

	bytes, err := base64.StdEncoding.DecodeString(prienctypt)

	prienctypt1, err := gorsa.PriKeyEncrypt(string(bytes), Pirvatekey)
	if err != nil {
		return err
	}

	pubdecrypt, err := gorsa.PublicDecrypt(prienctypt1, Pubkey)
	if err != nil {
		return err
	}

	toString := base64.StdEncoding.EncodeToString([]byte(pubdecrypt))

	pubdecrypt1, err := gorsa.PublicDecrypt(toString, Pubkey)
	if err != nil {
		return err
	}

	fmt.Println("string :", string(pubdecrypt1))

	return nil
}

// 公钥加密私钥解密
func applyPubEPriD() error {
	pubenctypt, err := gorsa.PublicEncrypt(`sadasiijkjkxkfjsafkjaskfjaskjfksajkfsajfklsajfqasfas
dsadsafsaffffjjsjvdskvxnzmvnvjiwoiwqoifcxskvnjfisajfiuwqiufoisfkjdsaklvjkdsajkfjsadkljfksajfjal
ccxzmvnbn<CXMqweupopwqpfidposicpuiwvknvmzxnmvn<MNC<MZXciiowquoieuqiwodkjsajfjdshvkjdsahfkjdsahf
ddqwoqo	wixxzcnvmzxsacsafasffqww`, Pubkey)
	if err != nil {
		return err
	}

	fmt.Println("加密信息:", pubenctypt)

	pridecrypt, err := gorsa.PriKeyDecrypt(`QLRJ9q3DWhngBRpYsOUfUH/57B6QBewSnGBRP3OwMcmcE9B3kDU19bc3CXlo5SwYak7VZLT9jkYsB474NVjtajBgGg1FbV2adgF+/7s8usfOk9cuWXJpWPAOeEtbsovlG10SbxxjJ71iha3vjeNRTn8Ms1mHR4OnxeOJcv0w0mKABVwNh6TYAeQs8yPka0Z9YH/QIyWh3e8l2eCYNwlj/a/Tp5xpf0xR0JQF2spagTX3MD9qyRImOJyBEMNxIjTsdHCyoDAMQtBYsKZcB6ykL13QwqjpxU3huesajGZrZKpQSinVNYyxNyZ14iv1uZ7jozOfETlxFTPchG6t97u1Fw==`, Pirvatekey)
	if err != nil {
		return err
	}

	fmt.Println("test :", pridecrypt)

	pridecrypt1, err := gorsa.PriKeyDecrypt(pridecrypt, Pirvatekey)
	if err != nil {
		return err
	}
	fmt.Println("test2 :", pridecrypt1)

	if string(pridecrypt) != `Hello world!` {
		return errors.New(`解密失败`)
	}
	return nil
}

// 公钥解密私钥加密
func applyPriEPubD() error {

	prienctypt, err := gorsa.PriKeyEncrypt(`hello world`, Pirvatekey)
	if err != nil {
		return err
	}

	pubdecrypt, err := gorsa.PublicDecrypt(prienctypt, Pubkey)
	if err != nil {
		return err
	}
	if string(pubdecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}
