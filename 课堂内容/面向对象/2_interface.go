package main

import "fmt"

type ts1 interface {
	Get() string
}

type myUrl struct {
	url string
}

func (my *myUrl) Get() string  {
	return my.url
}

func showInterface() ts1 {
	var urls myUrl
	urls.url = "one"

	return &urls
}

func main() {
	t1 := showInterface()
	fmt.Println(t1.Get())
}
