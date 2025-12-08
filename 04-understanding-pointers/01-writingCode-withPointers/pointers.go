package main

import "fmt"

func main() {
	age := 32 //Variavel normal

	fmt.Println("Age:", age)
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

//Esse 32 que é passado aqui é uma copia do 32 que definimos ali acima
func getAdultYears(age int) int {
	return age - 18
}
