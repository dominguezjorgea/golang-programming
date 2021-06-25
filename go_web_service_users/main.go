package main

import (
	"net/http"

	"github.com/go/webservice/controllers"
)

func main() {
	/*fmt.Println("Hello from a module gophers!")
	u := models.User{
		ID:        2,
		FirstName: "Jorge",
		LastName:  "Dominguez",
	}
	fmt.Println(u)*/
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
