package main

import (
	"fmt"
	"time"
)

type user struct { //criando a estrutura para agrupar as variaves sobre um único nome
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time //usando a função time, para definir o horário de criação
}

func main() {
	userfirstName := getUserData("Please, enter your first name: ")
	userlastName := getUserData("Please, enter your last name: ")
	userbirthdate := getUserData("Please, enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{ //Instanciando a estrutura
		firstName: userfirstName, //Passando atributo e valor, //podemos omitir o uso do "atributo:" no caso da ordem ser a mesma da estrutura exp: userfirsname,
		lastName:  userlastName,
		birthdate: userbirthdate,
		createdAt: time.Now(), //Instanciando atravez de função time
	}

	//... coloque algo incrível aqui

	outputUserDetails(userfirstName, userlastName, userbirthdate)
}

func outputUserDetails(firstName, lastName, birthdate string) {
	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
