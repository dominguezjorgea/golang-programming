package main

func main() {
	slice := []int{1, 2, 3}

	println("For with a postclause")
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}
}
