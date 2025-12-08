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
	firstName := getUserData("Please, enter your first name: ")
	lastName := getUserData("Please, enter your last name: ")
	birthdate := getUserData("Please, enter your birthdate (MM/DD/YYYY): ")

	//... coloque algo incrível aqui

	outputUserDetails(firstName, lastName, birthdate)
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
