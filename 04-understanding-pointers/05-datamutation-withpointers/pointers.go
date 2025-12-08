package main

import "fmt"

func main() {

	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer) //Não estamos mais armazenado nada, pois a função está alterando diretamente a variável
	fmt.Println(age)                //Variável alterada
}

//Alteramos a informação diretamente na váriavel, sem criar uma copia
//Precebesse, que em momento algum passamos a variável e sim seu endereço
func editAgeToAdultYears(age *int) {
	*age = *age - 18
	//Retirado o return, pois a operação está sendo feita diretamente na variavel
}
