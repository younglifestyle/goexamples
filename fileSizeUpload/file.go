package main

import (
	"fmt"
	"math"
)

func main() {
	var size = 5<<10<<10 + 5<<10<<10

	fmt.Println(5 << 10 << 10)

	fmt.Println(int(math.Ceil(float64(size) / (1 << 10 << 10))))
}
