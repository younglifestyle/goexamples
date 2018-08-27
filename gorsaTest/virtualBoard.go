package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"encoding/base64"
	"log"

	"io/ioutil"
	"os"

	"github.com/wenzhenxi/gorsa"
)

type License struct {
}

type VerifyLicenseData struct {
	Licenseid int    `json:"id" gorm:"column:licenseid"`
	Creator   int    `json:"c" gorm:"column:creator"`
	User      int    `json:"u" gorm:"column:user"`
	Index     int    `json:"i" gorm:"column:index"`
	Created   int64  `json:"cd" gorm:"column:created"`
	ValidDate int64  `json:"vd" gorm:"column:validdate"`
	DeviceId  string `json:"d" gorm:"column:deviceid"`
	Timestamp int64  `json:"t" gorm:"column:timestamp"`
	ValueA    int    `json:"va"`
	ValueB    int    `json:"vb"`
}

type CertDecryptData struct {
	DecryptCertStr VerifyLicenseData `json:"dc,omitempty"`
	HashMd5        string            `json:"h,omitempty"`
}

var jsonStr = []byte(`
{
  "dc":{
    "id":1,
    "c":4294967295,
    "u":123,
    "i":1,
    "cd": 1534305343,
    "vd": 20190801,
    "d": "testDeviceID",
    "t":  12314214,
    "va": 314212152,
    "vb": 214212153
  }
}
`)
var Pubkey = `-----BEGIN -----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDLftGiduO2ZS1m4dvQXYUuFak9
Ryo7lRrftsKxpv58+kY1UqgymbwPUW12/dXhSYEE7NNUzE9uB39qT4twALb7yIFO
QFmlsU0ymoiNCkDzUlECPABmeo5MjM5T+L4FEh53oRbgR/AotEQJw3/uIVGs0SFd
XI1rb4kX2r/ZmpbcVQIDAQAB
-----END -----
`

func CreateMd5Sum(sumData []byte) string {

	data := fmt.Sprintf("%x", md5.Sum(sumData))

	return data[len(data)-8 : len(data)]
}

func ReadFileAll(path string) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

func main() {
	//testModel := VerifyLicenseData{
	//	Licenseid: 1,
	//	Creator:   1,
	//	User:      123,
	//	Index:     1,
	//	Created:   20190801,
	//	ValidDate: 20190801,
	//	DeviceId:  "testDeviceID",
	//	Timestamp: 1232142142121323,
	//	ValueA:    314212152,
	//	ValueB:    214212153,
	//}
	//
	//bytes, _ := json.Marshal(testModel)
	//fmt.Println(string(bytes))

	tmp, _ := ReadFileAll("public.pem")
	tmp1, _ := ReadFileAll("priv.pem")
	Pubkey = string(tmp)
	PrivateKey := string(tmp1)

	all, _ := ReadFileAll("./board.json")

	tst := VerifyLicenseData{}
	json.Unmarshal(all, &tst)
	fmt.Println("\n file \n", string(all), tst)

	bytes, _ := json.Marshal(tst)

	certJsonData := CertDecryptData{
		DecryptCertStr: tst,
		HashMd5:        CreateMd5Sum(bytes),
	}

	data, _ := json.Marshal(certJsonData)

	fmt.Println("\n data \n", string(data))

	encryCall := func(needStr string) {
		pubenctypt, err := gorsa.PublicEncrypt(string(needStr), Pubkey)
		if err != nil {
			log.Println("ts1", err)
			return
		}

		databs, _ := base64.StdEncoding.DecodeString(pubenctypt)

		pubenctypt1, err := gorsa.PublicEncrypt(string(databs), Pubkey)
		if err != nil {
			log.Println("ts2", err)
			return
		}

		fmt.Println(pubenctypt1)
	}

	encryCall(string(data))

	all, _ = ReadFileAll("./offline.json")

	fmt.Println(" \n test \n", string(all))
	pubenctypt, err := gorsa.PriKeyEncrypt(string(all), PrivateKey)
	if err != nil {
		log.Println("ts1", err)
		return
	}

	fmt.Println(pubenctypt)
}
