package main

import (
	"fmt"

	"github.com/xtgo/set"
)

func main() {
	s := set.Strings([]string{"alpha", "gamma", "alpha"})
	fmt.Println("set:", s)

	s = set.StringsDo(set.Union, s, "beta")
	fmt.Println("set + beta:", s)

	fmt.Println(s, "contains any [alpha delta]:",
		set.StringsChk(set.IsInter, s, "alpha", "delta"))

	fmt.Println(s, "contains all [alpha delta]:",
		set.StringsChk(set.IsSuper, s, "alpha", "delta"))
}
