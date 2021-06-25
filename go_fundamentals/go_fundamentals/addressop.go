package main

import (
	"fmt"
)

func main() {
	firstName := "Jorge"
	fmt.Println(firstName)
	ptr := &firstName
	fmt.Println(ptr, *ptr)
	firstName = "Armando"
	fmt.Println(ptr, *ptr)
}
