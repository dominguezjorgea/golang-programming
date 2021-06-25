package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) error {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return errors.New("Something went wrong")
}
