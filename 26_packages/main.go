package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/Hydra-Ramesh/auth"
	"github.com/Hydra-Ramesh/model"
)

func main() {
	auth.LoginUserCredentials("Ramesh2004@", "secret")
	session := auth.GetSession()
	println("Session Status:", session)

	user := model.User{
		UserName: "Ramesh2004@",
		Password: "secret",
		Email:    "ramesh@example.com",
	}
	fmt.Printf("User Details: %+v\n", user)

	color.Red(user.UserName)
	color.Green(user.Email)
	color.Blue(user.Password)

}
