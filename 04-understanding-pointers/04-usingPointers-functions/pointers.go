package main

import "fmt"

func main() {

	age := 32

	var agePointer *int

	//Criando o ponteiro(local) da idade
	agePointer = &age

	fmt.Println("Age:", *agePointer)

	//Passando ponteiro para função
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

//Função pega o local(ponteiro) da variavel, evitando a criação de uma copia
//Nota, isso é irrelevante para números, numeros não ocupam quase nada de espaço na memória
func getAdultYears(age *int) int {
	return *age - 18 //Ao fazer o calculo, pegamos o valor desse ponteiro
}
