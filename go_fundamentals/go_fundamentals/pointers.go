package main

import (
	"fmt"
)

func main() {
	var secondName *string = new(string)
	*secondName = "Armando"
	fmt.Println(*secondName)
}
