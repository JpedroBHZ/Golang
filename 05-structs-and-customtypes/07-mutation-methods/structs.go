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

//Ponteiro é colocado aqui, para que seja feita a impressão do endereço e não da copia
func (u *user) outputUserDetails() { 
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

//A função não iria funcionar se não aceitasse o ponteiro, pois estariamos zerando a copia das variaveis
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
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

	appUser.outputUserDetails()
	appUser.clearUserName() //Mutação
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
