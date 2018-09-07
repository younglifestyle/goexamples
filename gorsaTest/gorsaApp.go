package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"encoding/base64"
	"strconv"
	"strings"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/toolkits/sys"
	"github.com/wenzhenxi/gorsa"
)

var (
	help    bool
	rsaData string
	method  string
	encry   bool // 加密
	decry   bool // 解密
	create  bool
	double  bool
)

func init() {
	flag.BoolVar(&help, "h", false, "[h] help 帮助")
	flag.BoolVar(&double, "dd", false, "两次加密/解密")
	flag.BoolVar(&create, "c", false, "[c] 生成需要key文件")
	flag.StringVar(&method, "m", "", `app -m, please input [pub]/[pri] 指定使用公钥或私钥,"pub/pri"`)
	flag.BoolVar(&encry, "e", false, "encrypt data, 加密数据")
	flag.BoolVar(&decry, "d", false, "decrypt data, 解密数据")
	flag.StringVar(&rsaData, "f", "", "app -f, Specify RSA Encrypt/Decrypt data 指定需要RSA加密/解密的数据")
}

/*
	本地方式
*/
func main() {
	flag.Parse()

	switch {
	case help:
		flag.Usage()
	case encry && rsaData != "":
		if method == "pub" {

			if double {
				str, err := PublicEncry(rsaData)
				if err != nil {
					fmt.Println("Public Encrypy :", err)
				}
				byteStr, _ := base64.StdEncoding.DecodeString(str)
				strs, err := PublicEncry(string(byteStr))
				if err != nil {
					fmt.Println("Public Encrypy :", err)
				}
				fmt.Println("result :", strs)

				return
			}

			_, err := PublicEncry(rsaData)
			if err != nil {
				fmt.Println("Public Encrypy :", err)
			}
		} else if method == "pri" {
			if double {
				str, err := PrivEncry(rsaData)
				if err != nil {
					fmt.Println("Public Encrypy :", err)
				}
				byteStr, _ := base64.StdEncoding.DecodeString(str)
				strs, err := PrivEncry(string(byteStr))
				if err != nil {
					fmt.Println("Public Encrypy :", err)
				}
				fmt.Println("result :", strs)
				return
			}

			_, err := PrivEncry(rsaData)
			if err != nil {
				fmt.Println("private Encrypy :", err)
			}
		} else {
			fmt.Println("./app -m, please input [pub]/[pri]")
			flag.Usage()
		}
	case decry && rsaData != "":
		if method == "pub" {
			if double {
				str, err := PublicDecry(rsaData)
				if err != nil {
					fmt.Println("Public Encrypy :", err)
				}
				byteStr := base64.StdEncoding.EncodeToString([]byte(str))

				strs, err := PublicDecry(string(byteStr))
				if err != nil {
					fmt.Println("Public Encrypy :", err)
				}
				fmt.Println("result :", strs)
				return
			}

			_, err := PublicDecry(rsaData)
			if err != nil {
				fmt.Println("Public Decrypt :", err)
			}
		} else if method == "pri" {

			if double {
				str, err := PrivDecry(rsaData)
				if err != nil {
					fmt.Println("Public Encrypy :", err)
				}
				byteStr := base64.StdEncoding.EncodeToString([]byte(str))

				strs, err := PrivDecry(string(byteStr))
				if err != nil {
					fmt.Println("Public Encrypy :", err)
				}
				fmt.Println("result :", strs)
				return
			}

			_, err := PrivDecry(rsaData)
			if err != nil {
				fmt.Println("private Decrypt :", err)
			}
		} else {
			fmt.Println("./app -m, please input [pub]/[pri]")
			flag.Usage()
		}
	case create:
		i := 0
		for {
			GenerateRsaKey("private.pem")
			modStr, privEStr, _ := GetKeyModulus("private.pem")
			i++
			fmt.Println("cycle :", i)

			if modStr == "" || privEStr == "" {
				continue
			} else {
				break
			}
		}
	default:
		flag.Usage()
	}
}

var PubKey = "public.pem"
var PrivKey = "private.pem"

// 公钥加密
func PublicEncry(needEncryData string) (pubenctypt string, err error) {

	pubkey, err := ReadFileAll(PubKey)
	if err != nil {
		log.Println("read pubkey is failed", err)
		return
	}
	pubenctypt, err = gorsa.PublicEncrypt(needEncryData, string(pubkey))
	if err != nil {
		log.Println("Encry is failed", err)
		return
	}

	fmt.Println("Result is :", string(pubenctypt))

	return
}

// 私钥加密
func PrivEncry(needEncryData string) (prienctypt string, err error) {

	privkey, err := ReadFileAll(PrivKey)
	if err != nil {
		log.Println("read privKey is failed")
		return
	}

	prienctypt, err = gorsa.PriKeyEncrypt(needEncryData, string(privkey))
	if err != nil {
		log.Println("Encry is failed")
		return
	}

	fmt.Println("Result is :", string(prienctypt))

	return
}

// 私钥解密
func PrivDecry(needDecryData string) (privDecryptRes string, err error) {

	privkey, err := ReadFileAll(PrivKey)
	if err != nil {
		log.Println("read decrypt data is failed")
		return
	}

	privDecryptRes, err = gorsa.PriKeyDecrypt(needDecryData,
		string(privkey))
	if err != nil {
		return
	}

	fmt.Println("Result is :", privDecryptRes)

	return
}

// 公钥解密
func PublicDecry(needDecryData string) (publicDecryptRes string, err error) {

	pubkey, err := ReadFileAll(PubKey)
	if err != nil {
		log.Println("read pubkey is failed", err)
		return
	}

	publicDecryptRes, err = gorsa.PublicDecrypt(needDecryData,
		string(pubkey))
	if err != nil {
		log.Println("PublicDecrypt is failed", err)
		return
	}

	fmt.Println("Result is :", publicDecryptRes)

	return
}

func ReadFileAll(path string) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

//生成 私钥和公钥文件
func GenerateRsaKey(pirvKeyPathName string) (privKey []byte, pubKey []byte, err error) {
	//生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	// 解析成[]byte
	privKey = pem.EncodeToMemory(block)

	file, err := os.Create(pirvKeyPathName)
	if err != nil {
		return nil, nil, err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, nil, err
	}

	//生成公钥文件
	//publicKey := &privateKey.PublicKey
	//defPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	//if err != nil {
	//	return nil, nil, err
	//}
	//block = &pem.Block{
	//	Type:  "RSA PUBLIC KEY",
	//	Bytes: defPkix,
	//}
	//
	//// 解析成[]byte
	//pubKey = pem.EncodeToMemory(block)
	//
	//file, err = os.Create(pubKeyPathName)
	//if err != nil {
	//	return nil, nil, err
	//}
	//err = pem.Encode(file, block)
	//if err != nil {
	//	return nil, nil, err
	//}

	return privKey, pubKey, nil
}

// openssl rsa -in private.pem -text -noout
func GetKeyModulus(pirvKeyPathName string) (modStr string, privEStr string, err error) {

	outStr, err := sys.CmdOut("openssl", "rsa",
		"-in", pirvKeyPathName, "-text", "-noout")
	if err != nil {
		log.Println("the command is not exist", err)
		return "", "", err
	}

	modIndex := strings.Index(outStr, "modulus")
	pubExIndex := strings.Index(outStr, "publicExponent")
	tmpStr := outStr[(modIndex + 8):(pubExIndex - 1)]
	tmpStr = strings.Replace(tmpStr, "\n", "", -1)
	tmpStr = strings.Replace(tmpStr, " ", "", -1)

	if tmpStr[:2] == "00" {
		//tmpStr = tmpStr[3:]
		return "", "", nil
	}

	outArr := strings.Split(tmpStr, ":")
	byteArr := make([]byte, len(outArr))
	for index, data := range outArr {
		i, _ := strconv.ParseUint(data, 16, 8)

		byteArr[index] = byte(i)
	}

	fmt.Println("modules :", byteArr)
	modStr = base64.StdEncoding.EncodeToString(byteArr)

	// privateExponent
	privE := strings.Index(outStr, "privateExponent")
	prime1Index := strings.Index(outStr, "prime1")
	tmpStr2 := outStr[(privE + 16):(prime1Index - 1)]
	tmpStr2 = strings.Replace(tmpStr2, "\n", "", -1)
	tmpStr2 = strings.Replace(tmpStr2, " ", "", -1)

	if tmpStr2[:2] == "00" {
		//tmpStr = tmpStr[3:]
		return "", "", nil
	}

	outArr2 := strings.Split(tmpStr2, ":")
	byteArr2 := make([]byte, len(outArr2))
	for index, data := range outArr2 {
		i, _ := strconv.ParseUint(data, 16, 8)

		byteArr2[index] = byte(i)
	}

	fmt.Println("privateExponent :", byteArr2)
	privEStr = base64.StdEncoding.EncodeToString(byteArr2)

	return
}

//func main() {
//
//	var needDecrypt string
//	var needEncrypt string
//	flag.StringVar(&needDecrypt, "d", "", "-d, need decrypt data")
//	flag.StringVar(&needEncrypt, "e", "", "-e, need Encrypt data")
//	flag.Parse()
//	if needDecrypt == "" || needEncrypt == "" {
//		fmt.Println("Please input need encrypt/decrypt data, ./app -d/-e xxxxx")
//		return
//	}
//
//	fi, _ := os.Open("public.pem")
//	defer fi.Close()
//	bytes, _ := ioutil.ReadAll(fi)
//
//	Pubkey = string(bytes)
//
//	pubdecrypt, err := gorsa.PublicDecrypt(needDecrypt, Pubkey)
//	if err != nil {
//		log.Println("error1 :", err)
//		return
//	}
//
//	toString := base64.StdEncoding.EncodeToString([]byte(pubdecrypt))
//	pubdecrypt1, err := gorsa.PublicDecrypt(toString,
//		Pubkey)
//	if err != nil {
//		log.Println("error2 :", err)
//		return
//	}
//
//	//fmt.Println("bytes :", pubdecrypt1)
//	var test map[string]interface{}
//	err = json.Unmarshal([]byte(pubdecrypt1), &test)
//	if err == nil {
//		fmt.Println()
//		for key, val := range test {
//			if key == "modulus" || key == "priv" {
//				strs, _ := base64.StdEncoding.DecodeString(val.(string))
//				fmt.Println(key, ":", strs)
//				fmt.Println()
//			} else {
//				fmt.Println(key, ":", val)
//				fmt.Println()
//			}
//		}
//		fmt.Println()
//	}
//}
