package main

import (
	"fmt"
	"github.com/pluralsight/webservice/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Peter",
		LastName: "Gallacher",
	}
	fmt.Println(u)
}