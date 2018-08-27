package main

import (
	"fmt"
	"strconv"
)

func main() {

	paramStr := make([]string, 5)

	paramStr[0] = "2131"
	paramStr[1] = "2132"
	paramStr[2] = "2133"
	paramStr[3] = "2134"
	paramStr[4] = "2135"
	for _, val := range paramStr {
		i, err := strconv.Atoi(val)
		fmt.Printf("%v %v \n", i, err)

	}
}
