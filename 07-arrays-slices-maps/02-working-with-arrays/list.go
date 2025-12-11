package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"} //Definindo o 1° item
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet" //Definindo o terceiro item

	fmt.Println(productNames) //Imprimindo o Array completo

	fmt.Println(prices[2]) //Lendo o terceiro item
}
