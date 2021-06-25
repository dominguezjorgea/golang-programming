package main

import "fmt"

func main() {

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u0 user
	fmt.Println(u0)
	u1 := user{
		ID:        1,
		FirstName: "Jorge",
		LastName:  "Dominguez",
	}
	fmt.Println(u1)

	u2 := user{ID: 2, FirstName: "Nidia", LastName: "Morales"}
	fmt.Println(u2)

	var u3 user
	u3.ID = 3
	u3.FirstName = "Juan"
	u3.LastName = "Dominguez"
	fmt.Println(u3)
}
