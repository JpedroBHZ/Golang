package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userfirstName := getUserData("Please, enter your first name: ")
	userlastName := getUserData("Please, enter your last name: ")
	userbirthdate := getUserData("Please, enter your birthdate (MM/DD/YYYY): ")

	//Instanciando
	var appUser *user.User

	//Construtor
	appUser, err := user.New(userfirstName, userlastName, userbirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
