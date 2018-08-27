package main

import (
	"fmt"

	"encoding/base64"

	"bytes"

	"github.com/dgkang/rsa/rsa"
)

var PubKey = "public.pem"
var PrivKey = "private.pem"

func main() {
	//file, err := os.Open("./test.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//b, e := ioutil.ReadAll(file)
	//if e != nil {
	//	log.Fatal(e)
	//}

	//b := []byte(`hello,world`)
	//
	//encrypt117 := divideCgoEncrypt117(b)
	//divideCgoEncrypt117(encrypt117)

	brsa := `gszjqdTZEAwkPP60XqyfirdZPoDU+mNPOxVF7AF2rUyV0VX3k2t44fvN7YU5sbuqJMmlXsXqPKf5nqgPZR3VuRPFeteaGHKACh7HTprff04rCDTj+SZcEj5vTM7QJ1JeVACo5zss1R++XS/jjxu8JR4MPSqqTFNaii3vZkLpjOw=`

	brsaB, _ := base64.StdEncoding.DecodeString(brsa)
	buf, e := rsa.PublicDecrypt(brsaB, PubKey, rsa.RSA_NO_PADDING)
	if e == nil {
		fmt.Printf("Decrypt: %s, len:%d \n", string(buf), len(buf))
		fmt.Println(buf[:], "\n" /*, string(buf[:116]), "   ", string(buf[116:])*/)
	} else {
		fmt.Printf("%s\n", e.Error())
		return
	}

	//zeroArr := make([]byte, 105)
	//strb := []byte(`hello,dsafaS`)
	//brsa1 := BytesCombine(buf[:11], zeroArr, strb)
	//
	//fmt.Println("2 len :", len(brsa1))
	//
	//brsa2, e := rsa.PrivateEncrypt(brsa1, "./private.pem", rsa.RSA_NO_PADDING)
	//if e != nil {
	//	fmt.Printf("%s\n", e.Error())
	//	return
	//}
	//fmt.Println(base64.StdEncoding.EncodeToString(brsa2))

	//ioutil.WriteFile("./public.rsa", brsa, os.ModePerm)
	//
	//buf, e := rsa.PublicDecrypt(brsa, "./public.pem", rsa.RSA_PKCS1_PADDING)
	//if e == nil {
	//	fmt.Printf("Decrypt: %s", string(buf))
	//} else {
	//	fmt.Printf("%s\n", e.Error())
	//	return
	//}
}

//BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func divideCgoEncrypt117(str []byte) []byte {

	var priEnc []byte

	lens := len(str)
	oneBlockSize := lens / 117

	for i := 0; i < oneBlockSize; i++ {

		prienctypt, err := rsa.PrivateEncrypt(str[i*117:(117+i*117)], PrivKey, rsa.RSA_NO_PADDING)
		if err != nil {
			fmt.Println("error 1", err)
			return nil
		}

		priEnc = BytesCombine(priEnc, prienctypt)
	}

	fmt.Println("byte :", len(str[(oneBlockSize*117):]))

	prienctypt, err := rsa.PrivateEncrypt(str[(oneBlockSize*117):], PrivKey, rsa.RSA_NO_PADDING)
	if err != nil {
		fmt.Println("error 2", err)
		return nil
	}

	priEnc = BytesCombine(priEnc, prienctypt)
	fmt.Println(base64.StdEncoding.EncodeToString(priEnc))

	return priEnc
}
