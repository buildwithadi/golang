package main

import (
	"24packages/auth"
	"24packages/user"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("Aditya", "password")

	session := auth.GetSession()
	fmt.Println(session)

	user2 := user.User{
		Email: "john@email.com",
		Name:  "John",
	}

	fmt.Println(user2.Email)

	color.Red("Program Ended")
}
