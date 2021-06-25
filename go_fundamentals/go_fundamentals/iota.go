package main

import (
	"fmt"
)

const (
	firt   = iota + 6
	second = 2 << iota
)

const (
	third = iota
	fourth
)

func main() {
	//iota increments automatically each time is being used
	fmt.Println(firt, second, third)
}
