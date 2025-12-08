package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userfirstName := getUserData("Please, enter your first name: ")
	userlastName := getUserData("Please, enter your last name: ")
	userbirthdate := getUserData("Please, enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userfirstName,
		lastName:  userlastName,
		birthdate: userbirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(appUser) //Passando estrutura para função
}

func outputUserDetails(u user) { //Aceitando variavel do tipo user(estrutura instanciada)
	fmt.Println(u.firstName, u.lastName, u.birthdate) //Passando atributos da estrutura
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
