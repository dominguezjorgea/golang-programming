package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	u2 := User{
		ID:        2,
		FirstName: "Ford",
		LastName:  "Prefect",
	}

	//u3 := *&u1
	if u1 == u2 {
		//if u1 == u3 {
		println("Same user!")
		//} else if u1.FirstName == u3.FirstName {
	} else if u1.FirstName == u2.FirstName {
		println("Similar user.")
	} else {
		println("Different user!")
	}
}
