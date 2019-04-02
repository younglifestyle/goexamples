package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.Rename("./tmp", "./tmp1")

	if err != nil {
		fmt.Println(err)
		return
	}
}
