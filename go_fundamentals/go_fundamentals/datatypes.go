package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from a module, Gophers!!!")
	var i int
	i = 42
	fmt.Println(i)
	//two types of float: float32 and float64
	var f float32 = 3.14
	fmt.Println(f)
	firstName := "Jorge"
	fmt.Println(firstName)
	b := true
	fmt.Println(b)
	//built-in funtion complex
	c := complex(3, 4)
	fmt.Println(c)
	//Multiple assignments
	r, il := real(c), imag(c)
	fmt.Println(r, il)
}
