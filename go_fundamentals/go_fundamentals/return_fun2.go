package main

import "fmt"

func main() {
	port := 3000
	//port, err := startWebServer(port)
	//_ ignores the first value returned
	_, err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return port, nil
}
