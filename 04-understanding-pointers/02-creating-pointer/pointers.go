package main

import "fmt"

func main() {
	age := 32

	//Parando o mouso sobre a variavel vemos o tipo dela, e percebemos q variaveis com * viram ponteiros
	var agePointer *int

	//Colocando e comercial antes da variavel ela vira um ponteiro (referenciamos o endereço do valor)
	agePointer = &age

	fmt.Println("Age:", agePointer)

	//adultYears := getAdultYears(age)
	///fmt.Println(adultYears)
}

//func getAdultYears(age int) int {
//	return age - 18
//}
