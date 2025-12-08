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

//colocando u user(argumento do receptor), damos acesso a struct para a func, transformando-a em um metodo
//Simplificando, aceitamos uma variavel x do tipo user, e usamos ela para definir quais atributos vamos usar

func (u user) outputUserDetails() { 
	fmt.Println(u.firstName, u.lastName, u.birthdate)
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

	//Agr passamos o usuario e invocamos o metodo
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
