package main

import "fmt"

func main() {
	port := 3000
	isStarted := startWebServer(port)
	fmt.Println(isStarted)
}

func startWebServer(port int) bool {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return true
}
