package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	slice := arr[:]
	arr[1] = 42
	slice[2] = 27 //Kind of a pointer

	fmt.Println(arr, slice)
}
