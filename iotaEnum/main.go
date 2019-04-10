package main

import (
	"fmt"
)

const (
	// B byte
	B = (int64)(1 << (10 * iota))
	// KB kilobyte
	KB
	// MB megabyte
	MB
	// GB gigabyte
	GB
	// TB terabyte
	TB
	// PB petabyte
	PB
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(B)
	fmt.Println(KB, MB)

	var d Direction = South
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
