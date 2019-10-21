package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

func main() {
	val, err := ioutil.ReadFile("test.toml")
	if err != nil {
		fmt.Println("read file error, ",err)
		return
	}

	var tmp interface{}
	md, err := toml.Decode(string(val), &tmp)
	if err != nil {
		fmt.Println("decode error, ",err)
		return
	}

	fmt.Println(md)
	fmt.Println(tmp)

	bytes, _ := json.Marshal(tmp)
	fmt.Println(string(bytes))
}
