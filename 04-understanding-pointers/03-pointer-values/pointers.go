package main

import "fmt"

func main() {

	//Todas as variaveis quando definidas vem 0 ou "" e obtem endereços na memória, menos as definidas "nil"
	age := 32

	var agePointer *int

	//Colocando e comercial antes da variavel ela vira um ponteiro (referenciamos o endereço do valor)
	agePointer = &age

	//Colocando asterisco referenciamos de volta o valor do ponteiro
	fmt.Println("Age:", *agePointer)

	//adultYears := getAdultYears(age)
	//fmt.Println(adultYears)
}

//func getAdultYears(age int) int {
//	return age - 18
//}
